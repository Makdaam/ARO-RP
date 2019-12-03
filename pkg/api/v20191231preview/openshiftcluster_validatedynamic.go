package v20191231preview

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/apparentlymart/go-cidr/cidr"

	"github.com/jim-minter/rp/pkg/api"
	"github.com/jim-minter/rp/pkg/util/subnet"
)

type dynamicValidator struct {
	subnets subnet.Manager

	oc *api.OpenShiftCluster
	r  azure.Resource
}

// validateOpenShiftClusterDynamic validates an OpenShift cluster
func validateOpenShiftClusterDynamic(ctx context.Context, fpAuthorizer autorest.Authorizer, oc *api.OpenShiftCluster) error {
	r, err := azure.ParseResourceID(oc.ID)
	if err != nil {
		return err
	}

	v := &dynamicValidator{
		oc: oc,
		r:  r,
	}

	spAuthorizer, err := v.validateServicePrincipalProfile()
	if err != nil {
		return err
	}

	vnetID, _, err := subnet.Split(oc.Properties.MasterProfile.SubnetID)
	if err != nil {
		return err
	}

	ok, err := v.hasOwnerPermissionOnVnet(ctx, fpAuthorizer, vnetID)
	if err != nil {
		return err
	}
	if !ok {
		return api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidResourceProviderPermissions, "", "The resource provider does not have Owner permission on vnet '%s'.", vnetID)
	}

	ok, err = v.hasOwnerPermissionOnVnet(ctx, spAuthorizer, vnetID)
	if err != nil {
		return err
	}
	if !ok {
		return api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidServicePrincipalPermissions, "", "The provided service principal does not have Owner permission on vnet '%s'.", vnetID)
	}

	v.subnets = subnet.NewManager(r.SubscriptionID, spAuthorizer)

	return v.validateSubnets(ctx)
}

func (dv *dynamicValidator) validateServicePrincipalProfile() (autorest.Authorizer, error) {
	spp := &dv.oc.Properties.ServicePrincipalProfile
	conf := auth.NewClientCredentialsConfig(spp.ClientID, spp.ClientSecret, spp.TenantID)

	token, err := conf.ServicePrincipalToken()
	if err != nil {
		return nil, err
	}

	err = token.EnsureFresh()
	if err != nil {
		return nil, api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidServicePrincipalCredentials, "properties.servicePrincipalProfile", "The provided service principal credentials are invalid.")
	}

	return conf.Authorizer()
}

func (dv *dynamicValidator) hasOwnerPermissionOnVnet(ctx context.Context, authorizer autorest.Authorizer, vnetID string) (bool, error) {
	vnetr, err := azure.ParseResourceID(vnetID)
	if err != nil {
		return false, err
	}

	permissions := authorization.NewPermissionsClient(dv.r.SubscriptionID)
	permissions.Authorizer = authorizer

	page, err := permissions.ListForResource(ctx, vnetr.ResourceGroup, vnetr.Provider, vnetr.ResourceType, "", vnetr.ResourceName)
	if err != nil {
		if err, ok := err.(autorest.DetailedError); ok {
			if err.StatusCode == http.StatusNotFound {
				return false, api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidLinkedVNet, "properties.masterProfile.subnetId", "The provided master VM subnet '%s' could not be found.", dv.oc.Properties.MasterProfile.SubnetID)
			}
		}

		return false, err
	}

	for {
		for _, p := range page.Values() {
			for _, a := range *p.Actions {
				if a == "*" {
					return true, nil
				}
			}
		}

		err = page.Next()
		if err != nil {
			return false, err
		}

		if !page.NotDone() {
			break
		}
	}

	return false, nil
}

func (dv *dynamicValidator) validateSubnets(ctx context.Context) error {
	master, err := dv.validateSubnet(ctx, "properties.masterProfile.subnetId", "master", dv.oc.Properties.MasterProfile.SubnetID)
	if err != nil {
		return err
	}

	worker, err := dv.validateSubnet(ctx, `properties.workerProfiles["worker"].subnetId`, "worker", dv.oc.Properties.WorkerProfiles[0].SubnetID)
	if err != nil {
		return err
	}

	_, pod, err := net.ParseCIDR(dv.oc.Properties.NetworkProfile.PodCIDR)
	if err != nil {
		return err
	}

	_, service, err := net.ParseCIDR(dv.oc.Properties.NetworkProfile.ServiceCIDR)
	if err != nil {
		return err
	}

	err = cidr.VerifyNoOverlap([]*net.IPNet{master, worker, pod, service}, &net.IPNet{IP: net.IPv4zero, Mask: net.IPMask(net.IPv4zero)})
	if err != nil {
		return api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidLinkedVNet, "", "The provided CIDRs must not overlap: '%s'.", err)
	}

	return nil
}

func (dv *dynamicValidator) validateSubnet(ctx context.Context, path, typ, subnetID string) (*net.IPNet, error) {
	s, err := dv.subnets.Get(ctx, subnetID)
	if err != nil {
		if err, ok := err.(autorest.DetailedError); ok {
			if err.StatusCode == http.StatusNotFound {
				return nil, api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidLinkedVNet, path, "The provided "+typ+" VM subnet '%s' could not be found.", subnetID)
			}
		}
		return nil, err
	}

	if dv.oc.Properties.ProvisioningState == api.ProvisioningStateCreating {
		if s.SubnetPropertiesFormat != nil &&
			s.SubnetPropertiesFormat.NetworkSecurityGroup != nil {
			return nil, api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidLinkedVNet, path, "The provided "+typ+" VM subnet '%s' is invalid: must not have a network security group attached.", subnetID)
		}

	} else {
		nsgID, err := subnet.NetworkSecurityGroupID(dv.oc, *s.ID)
		if err != nil {
			return nil, err
		}

		if s.SubnetPropertiesFormat == nil ||
			s.SubnetPropertiesFormat.NetworkSecurityGroup == nil ||
			!strings.EqualFold(*s.SubnetPropertiesFormat.NetworkSecurityGroup.ID, nsgID) {
			return nil, api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidLinkedVNet, path, "The provided "+typ+" VM subnet '%s' is invalid: must have network security group '%s' attached.", subnetID, nsgID)
		}
	}

	_, net, err := net.ParseCIDR(*s.AddressPrefix)
	if err != nil {
		return nil, err
	}
	{
		ones, _ := net.Mask.Size()
		if ones > 27 {
			return nil, api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeInvalidLinkedVNet, path, "The provided "+typ+" VM subnet '%s' is invalid: must be /27 or larger.", subnetID)
		}
	}

	return net, nil
}