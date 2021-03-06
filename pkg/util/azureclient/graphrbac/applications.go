package graphrbac

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

//go:generate go run ../../../../vendor/github.com/golang/mock/mockgen -destination=../../../util/mocks/azureclient/$GOPACKAGE/$GOPACKAGE.go github.com/Azure/ARO-RP/pkg/util/azureclient/$GOPACKAGE ApplicationsClient
//go:generate go run ../../../../vendor/golang.org/x/tools/cmd/goimports -local=github.com/Azure/ARO-RP -e -w ../../../util/mocks/azureclient/$GOPACKAGE/$GOPACKAGE.go

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
)

// ApplicationsClient is a minimal interface for azure ApplicationsClient
type ApplicationsClient interface {
	GetServicePrincipalsIDByAppID(ctx context.Context, applicationID string) (result graphrbac.ServicePrincipalObjectResult, err error)
}

type applicationsClient struct {
	graphrbac.ApplicationsClient
}

var _ ApplicationsClient = &applicationsClient{}

// NewApplicationsClient creates a new ApplicationsClient
func NewApplicationsClient(tenantID string, authorizer autorest.Authorizer) ApplicationsClient {
	client := graphrbac.NewApplicationsClient(tenantID)
	client.Authorizer = authorizer

	return &applicationsClient{
		ApplicationsClient: client,
	}
}
