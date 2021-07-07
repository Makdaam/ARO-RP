// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// deploy/staticresources/aro.openshift.io_clusters.yaml
// deploy/staticresources/master/deployment.yaml
// deploy/staticresources/master/rolebinding.yaml
// deploy/staticresources/master/service.yaml
// deploy/staticresources/master/serviceaccount.yaml
// deploy/staticresources/namespace.yaml
// deploy/staticresources/worker/deployment.yaml
// deploy/staticresources/worker/role.yaml
// deploy/staticresources/worker/rolebinding.yaml
// deploy/staticresources/worker/serviceaccount.yaml
package deploy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _aroOpenshiftIo_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x4d\x73\xe3\xc8\xcd\xbe\xeb\x57\xa0\xfc\x1e\x7c\x78\x2d\x79\xa6\x52\xa9\x4a\x74\xf3\xda\x3b\x13\xd5\xee\x78\x5c\xb6\x6b\xf6\xb0\xde\x03\xd8\x84\x28\xc4\xcd\x6e\x2e\xba\x29\x5b\x93\xca\x7f\x4f\xa1\x49\xea\xcb\xa4\x6c\x8d\x2b\x97\xf0\x32\xa3\xfe\x78\x00\x3c\x40\xa3\x81\xf6\x68\x3c\x1e\x8f\xb0\xe2\x6f\x24\x81\xbd\x9b\x02\x56\x4c\xcf\x91\x9c\xfe\x0a\x93\xc7\xbf\x85\x09\xfb\xf3\xe5\xc7\xd1\x23\xbb\x7c\x0a\x97\x75\x88\xbe\xbc\xa5\xe0\x6b\x31\x74\x45\x73\x76\x1c\xd9\xbb\x51\x49\x11\x73\x8c\x38\x1d\x01\xa0\x73\x3e\xa2\x0e\x07\xfd\x09\x60\xbc\x8b\xe2\xad\x25\x19\x17\xe4\x26\x8f\x75\x46\x59\xcd\x36\x27\x49\xe0\x9d\xe8\xe5\x87\xc9\x5f\x27\x1f\x46\x00\x46\x28\x6d\xbf\xe7\x92\x42\xc4\xb2\x9a\x82\xab\xad\x1d\x01\x38\x2c\x69\x0a\xc6\xd6\x21\x92\x84\x09\x8a\x9f\xf8\x8a\x5c\x58\xf0\x3c\x4e\xd8\x8f\x42\x45\x46\x65\x16\xe2\xeb\x6a\x0a\x2f\xe6\x1b\x84\x56\xad\xd6\xa4\x06\x2c\x8d\x58\x0e\xf1\x97\xed\xd1\x5f\x39\xc4\x34\x53\xd9\x5a\xd0\x6e\x44\xa7\xc1\xc0\xae\xa8\x2d\xca\x7a\x78\x04\x10\x8c\xaf\x68\x1b\xb5\x35\x2f\xc9\x1c\xb7\x06\x2c\x3f\xa2\xad\x16\xf8\xb1\x41\x31\x0b\x2a\xb1\x51\x09\x40\xd5\xbd\xb8\x99\x7d\xfb\xcb\xdd\xce\x30\x40\x4e\xc1\x08\x57\x31\x51\xd5\xc2\x03\x07\x88\x0b\x82\x66\x2d\xcc\xbd\xa4\x9f\x9d\x92\x70\x71\x33\x5b\xef\xaf\xc4\x57\x24\x91\x3b\xeb\x9b\x6f\xcb\xf5\x5b\xa3\x7b\xd2\x4e\x55\xa1\x66\x15\xe4\xea\x73\x6a\xc4\xb6\xa6\x51\xde\xda\x00\x7e\x0e\x71\xc1\x01\x84\x2a\xa1\x40\xae\x89\x82\x1d\x60\xd0\x45\xe8\xc0\x67\xff\x24\x13\x27\x70\x47\xa2\x30\x10\x16\xbe\xb6\xb9\x86\xca\x92\x24\x82\x90\xf1\x85\xe3\xef\x6b\xec\x00\xd1\x27\xa1\x16\x23\xb5\x4e\xd9\x7c\xec\x22\x89\x43\x0b\x4b\xb4\x35\x9d\x01\xba\x1c\x4a\x5c\x81\x90\x4a\x81\xda\x6d\xe1\xa5\x25\x61\x02\x5f\xbc\x10\xb0\x9b\xfb\x29\x2c\x62\xac\xc2\xf4\xfc\xbc\xe0\xd8\x85\xbc\xf1\x65\x59\x3b\x8e\xab\xf3\x14\xbd\x9c\xd5\xd1\x4b\x38\xcf\x69\x49\xf6\x3c\x70\x31\x46\x31\x0b\x8e\x64\x62\x2d\x74\x8e\x15\x8f\x93\xea\x2e\x85\xfd\xa4\xcc\xff\x4f\xda\x43\x12\x4e\x77\x74\x8d\x2b\x0d\x8f\x10\x85\x5d\xb1\x35\x91\x62\xf1\x80\x07\x34\x2a\xd5\xdb\xd8\x6e\x6d\xac\xd8\x10\xad\x43\xca\xce\xed\xcf\x77\xf7\xd0\x89\x4e\xce\xd8\x67\x3f\xf1\xbe\xd9\x18\x36\x2e\x50\xc2\xd8\xcd\x49\x1a\x27\xce\xc5\x97\x09\x93\x5c\x5e\x79\x76\xb1\x8d\x2d\x26\xb7\x4f\x7f\xa8\xb3\x92\xa3\xfa\xfd\xcf\x9a\x42\x54\x5f\x4d\xe0\x32\xe5\x01\xc8\x08\xea\x2a\xc7\x48\xf9\x04\x66\x0e\x2e\xb1\x24\x7b\x89\x81\xfe\xeb\x0e\x50\xa6\xc3\x58\x89\x7d\x9b\x0b\xb6\x53\xd8\xfe\xe2\x86\xb5\xad\x89\x2e\xd1\x0c\xf8\xab\x3d\x9f\x77\x15\x99\x9d\x13\x93\x53\x60\xd1\x98\x8e\x18\x49\x4f\xc2\x76\xf6\xe9\xbe\xfe\x93\xaa\x1f\x1a\xb9\xf2\x25\xb2\xdb\x9f\x18\x34\x0a\x9a\x33\x3e\x73\x71\x76\x73\xdc\xa6\x2d\x76\x7b\x33\xc4\x66\xbf\x1e\xbe\x62\xcf\x06\x00\xfc\xfe\xb3\x5b\xb2\x78\x57\x92\x8b\x47\x89\xce\xd0\x39\x92\x97\x5b\x76\x18\xfe\x29\x2d\x5a\x93\xcb\x73\xc0\x6e\xac\x4d\x25\x19\xe9\xff\x9e\x5c\x97\x38\x4c\xba\xbb\x5e\xe8\x79\x88\x6f\x68\x2f\xaf\x5e\x0b\x5e\x23\x60\x30\x76\x92\x2d\xc7\xbb\x71\x4e\xa8\xce\xe8\x51\x72\x87\x98\x4f\xed\xb2\x9d\xd8\xbb\xb8\xfd\xaa\x37\x8b\x60\xf4\xd2\x01\x41\xa1\x99\xf4\x48\x36\x34\x8f\x3a\xc3\x96\x2e\x2c\x49\xfc\x8d\xb2\x85\xf7\x8f\x87\xb8\xc9\xbc\xb7\x84\xfb\x77\xc0\x0e\xd4\x4f\x03\xfe\x3e\x0a\xe4\xea\xfa\xee\x0b\x86\x3f\xdf\x89\xf2\x99\x1c\x2d\xf1\x57\x5f\x14\xec\x8a\x77\x62\x7d\xf1\x8e\xa3\x57\x47\x5e\x7a\x37\xe7\xf7\xc2\x5d\xdf\x7d\xee\xf5\xc8\x31\x10\x3e\xa7\x2b\x41\x7e\x3f\xdb\x37\xb5\xb5\x77\x64\x84\x0e\x9e\x8b\x37\x00\xdd\xfa\x3a\xd2\x27\x7e\x7e\x27\xcc\x6f\x5e\x1e\x51\x7c\xed\xf2\x70\xb9\x2e\x36\x7f\x04\xf3\xc0\x91\x2d\x0e\xc7\xc6\xab\x59\x64\xce\xc5\x60\x22\x4d\x00\x18\xb5\x88\x99\xc2\xe9\xef\x1f\xc6\x7f\xff\xe3\xff\x27\xcd\x3f\xa7\x07\xac\xe8\xcd\x13\xfa\x95\xeb\xd8\xfb\x7c\x79\x77\x61\x8c\xaf\x87\x32\x18\xb9\xba\xec\x9f\x19\x6b\xd6\xe8\xea\x60\x5f\x84\xd9\xf5\xfd\x9b\xd6\xdd\xdc\x7e\xbd\x7a\xd3\xc2\x77\x1b\x76\xf0\x82\x79\xcd\xb8\x2b\xc6\xc2\xf9\x10\xd9\x84\x1b\xf1\xf9\xc0\xaa\xfb\x97\xb5\x66\x37\x75\x89\x9f\x90\x65\x8e\xcf\xef\xb6\xe3\x5a\x7b\x92\x0a\x0d\xfd\x0f\xb8\xe8\xc0\x01\x62\x37\x17\x9c\xe5\x47\x5d\x7a\xec\x0a\xa1\x10\x8e\x2c\x5e\x9a\x76\x80\xe2\xe5\x82\xcc\x63\x5f\x22\x38\x7c\x58\x6b\xb1\x03\xa9\x96\x23\x95\x03\x53\xaf\xba\xbc\x5b\x80\x22\xb8\x3a\x86\x37\xeb\x4d\xea\xa3\x8e\xa2\xa0\xeb\x01\xfa\xf8\xde\xa9\x17\xba\x66\x7e\x76\xd5\x75\x93\x17\xdf\xb5\x3a\xd8\x00\x34\x6d\x1d\x6d\x35\xb9\x6f\xd6\x62\xe9\x28\x1e\xe5\xf1\xa1\x72\x3b\x62\xac\xc3\x1b\x0a\xee\xb4\x6e\xa7\xe4\xf6\x59\xd0\xfe\xe6\x87\x6b\x6e\xe3\x5d\xce\x5b\x8f\x19\xc3\x2a\xac\x17\xb6\x8d\x1a\xc5\x24\xad\x1b\x06\x76\x21\xa2\x33\x14\x26\x2f\x80\x06\xe3\x6a\x47\xc2\xc9\x06\x6b\xd3\xbf\x35\xcd\xb4\xda\x98\x82\x64\xa7\xbd\x3e\xed\x3f\xc0\x89\x89\xc9\xb6\xc2\x28\xa4\xbb\xd6\x2f\x3f\x50\x92\x59\xa0\xe3\x50\xa6\xb3\xe4\x72\xca\xb5\x8c\xd6\x5e\x2e\x50\x7f\xae\x7c\x5a\x90\x6b\x7b\x9c\x88\x6c\xc3\x5a\x91\x8d\x6a\x2a\x45\x5b\x42\x84\x4a\xd8\x0b\xc3\xa3\xd3\xfa\xdc\x0b\x3c\x69\xc5\xde\x0b\x9b\xd6\x57\x95\x5d\xa9\x7c\xb4\x76\xc3\x62\x12\x00\x05\x2f\xc9\x81\xb6\xc7\x13\x78\x70\xdb\x36\x35\x6d\x40\x2f\x68\x46\x80\x79\x6b\x13\x3d\x57\x96\x0d\x47\xbb\x6a\x1e\x1f\x56\x5b\xb1\x00\x71\x81\x51\x4d\x96\x90\x9e\x14\x8c\x2f\x2b\xef\x94\xf5\x5e\x58\x93\x68\xcc\x7c\x1d\x41\x30\x2e\x52\x23\x8d\x2e\x75\xc5\x2c\x4d\x87\xee\x03\xed\xe0\x27\x4e\x53\xd3\x2d\x03\xbc\xa6\x36\xdc\x27\xb4\x2d\x2e\xc3\x04\xbe\x3a\x43\x6d\xa4\xe7\x67\x89\xf9\x92\xd0\xa9\x98\x44\xcc\x9a\x89\x01\x55\x1d\xb4\xdd\xb9\x3a\xba\xa0\x1c\x50\x32\x8e\x82\xc2\x76\x05\x63\x60\x9d\x33\xbe\xa4\x00\x15\x4a\xec\x72\xc0\xc5\xcd\x2c\xbd\xae\xf4\x82\x2e\xb0\x39\x72\x01\x4b\x82\x0c\xcd\xe3\x13\x4a\x1e\xc6\x89\xba\xb9\x97\xe6\x97\x72\x88\x91\x33\xb6\x1c\x13\xe5\x86\xc4\xa9\x33\x7b\x21\xd1\xad\x5a\xe3\xf7\xb4\x98\x9c\xf4\xac\x3f\x9c\xd6\x01\x2c\x86\x78\x2f\xe8\x02\x77\xcf\x8b\x43\xb9\x7c\xee\xa5\xc4\x38\x85\x1c\x23\x8d\x23\x97\xf4\xa3\x39\xbf\xa4\x10\xb0\x18\x94\xf3\xea\x7e\x21\x0c\x43\x15\xe3\x50\x02\xba\x4d\x7b\x34\x0b\xed\x1d\x5e\x04\xef\x68\xfc\xe4\x25\x3f\xdb\x3c\xc4\x0c\x40\xc3\xde\x2b\xde\xfa\x16\xc0\x48\x85\x97\x95\xfe\x36\x58\x07\x5a\x4f\xd4\x22\xe4\x62\x9b\xab\x5f\xe6\xb8\xee\x9b\xc5\x1e\xcd\x34\xad\x00\xbb\x14\x0f\xac\x98\x75\xac\xea\x78\x06\xa1\x36\x0b\xc0\x90\xf4\xb6\xec\x86\x95\x7d\xac\x33\x32\xd1\x42\xa1\x59\xb7\xdd\xac\x71\xc7\x0e\x42\x5d\x96\x28\xfc\x3d\x1d\x0d\xd3\xa8\xd9\xe6\x8f\x64\xc0\xa0\xae\xaf\x3a\xa7\xef\x5a\x3a\x62\x7b\x5a\xf0\x16\xcf\x6e\x12\xff\xfd\xaa\xa2\xee\x9e\xd6\xed\x6b\xf2\xd7\x37\xc3\xd0\xe1\xd4\x4f\x37\xae\x2a\x36\x68\xed\x4a\x53\x44\x17\x02\x39\x68\x4c\x68\x22\x0e\x0b\x2f\x11\xaa\x85\xa4\x17\xba\xed\x84\x3a\x08\x9a\xde\xd9\xba\x47\x17\x76\x39\x6b\x84\xb4\xb7\x2d\x37\x57\xc2\xc3\x09\x66\x4e\x4f\x94\x1d\x47\xa9\xe9\xe1\x04\x2a\x6f\x51\x38\xae\x86\xc3\xe4\x93\x17\xa0\x67\x2c\x2b\x4b\x67\xc0\xfb\x56\x76\x72\x42\x73\xef\xa0\x02\xb2\x59\x35\x91\xb5\x44\xcb\xf9\xd9\xb0\xc2\x49\x23\x0e\x90\xd6\x3d\x9c\x80\xc1\x90\x48\xad\xc4\x67\x98\xe9\x55\xb3\xd0\x8b\x4a\xca\x33\x08\x7e\x57\xf0\x20\xe8\xe6\xd1\x49\xd9\xa5\x1c\x1e\x4e\x66\xae\x15\xd0\x9b\xab\xe0\xf5\x08\x69\x2e\x0e\xea\xa9\x9f\xb4\x52\x6f\x82\xaf\x77\x4a\x71\x7b\x26\x0e\xd4\x98\x87\x8a\xd3\xee\xe1\xe8\x95\x67\xc0\x81\x32\x34\xff\x07\xc6\x5f\x68\x15\x6e\x9a\x5c\xf2\x72\xf7\x60\xed\xf3\x86\x3e\xe3\xa5\xba\xbd\x36\xbe\x18\x6c\x4a\xc2\x29\x68\x34\x36\x03\xd1\x8b\xa6\xe9\xad\x91\x3a\x5b\x3f\xe1\x77\xda\xb5\xe7\x1d\xfe\xf5\xef\xd1\xe6\xe8\xa3\x31\x54\x45\xca\xaf\xf7\xff\xb2\x74\xd2\xb8\xbd\xfb\xd3\x51\xfa\xb9\x55\x4d\xc2\xef\x7f\x8c\x1a\xc1\x94\x7f\xeb\xfe\x48\xa4\x83\xff\x09\x00\x00\xff\xff\x80\x35\xa6\x85\x94\x1b\x00\x00")

func aroOpenshiftIo_clustersYamlBytes() ([]byte, error) {
	return bindataRead(
		_aroOpenshiftIo_clustersYaml,
		"aro.openshift.io_clusters.yaml",
	)
}

func aroOpenshiftIo_clustersYaml() (*asset, error) {
	bytes, err := aroOpenshiftIo_clustersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aro.openshift.io_clusters.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\xcd\x6e\xdb\x30\x0c\xbe\xfb\x29\x88\xde\xdd\xa4\x87\x01\x85\x6e\xc5\x1a\xf4\xb2\x15\xc5\xba\xee\xce\xc8\x6c\x2c\x54\x12\x35\x92\x0e\xea\x3e\xfd\x20\x24\x71\xec\x15\xc8\x0e\xd3\xc9\x20\x3f\x7e\x3f\x94\x8c\x25\xfc\x22\xd1\xc0\xd9\x01\x96\xa2\xab\xfd\x4d\xf3\x16\x72\xe7\xe0\x9e\x4a\xe4\x31\x51\xb6\x26\x91\x61\x87\x86\xae\x01\x88\xb8\xa5\xa8\xf5\x0b\xea\x80\x03\x14\x6e\xb9\x90\xa0\xb1\xb4\x09\xd5\x48\x1a\x80\x8c\x89\x2e\xf5\xb4\xa0\x27\x07\x5c\x28\x6b\x1f\x5e\xad\xc5\x8f\x41\x68\x02\x37\x5a\xc8\x57\x11\xa1\x12\x83\x47\x75\x70\xd3\x00\x28\x45\xf2\xc6\x72\x90\x4f\x68\xbe\xff\x36\xf3\x73\xd1\x91\x9a\xa0\xd1\x6e\x3c\x40\x85\x63\x0c\x79\xf7\x52\x3a\x34\x3a\x4d\x27\x7c\x7f\x1e\x64\x47\x07\xb1\x63\xe5\x25\xe3\x1e\x43\xc4\x6d\x24\x07\xeb\x06\xc0\x28\x95\x38\x4d\xcd\x77\x53\x4f\x5c\xf8\xb9\xe8\x08\xe0\x94\xb2\x1e\xcf\xd9\x30\x64\x92\x69\xb8\x05\xcf\x29\x61\xee\xce\x6c\x6d\xa5\x3a\x73\xcb\x4e\xe7\xbd\x69\x7b\xe7\xd2\x4c\xac\x9e\x90\xb0\xc6\x7b\xd8\x3c\x6e\x7e\xdc\xfd\xdc\xdc\x4f\x8d\xcf\xf7\x35\xb5\x84\x94\x07\xf1\x34\x93\x02\x88\x21\x05\x5b\x54\x00\x7c\x19\x1c\x7c\x59\xa7\x45\x31\x51\x62\x19\x6b\x7d\xfd\x3d\xcc\x3a\x42\xbf\x07\xd2\xff\xa0\x28\x2c\xb6\x48\x3f\x2d\xf0\x89\xc5\x1c\xdc\xae\x6f\xd7\x33\x9a\x43\xc0\xde\xac\x34\xe7\x0c\x7b\xca\xa4\xfa\x24\xbc\xa5\xb9\x8f\x8a\x7a\x20\x5b\x5a\x2b\x68\xbd\x83\x55\x4f\x18\xad\xff\x58\x09\x61\x37\x2e\x01\x7f\xcb\x66\xee\xe8\x79\xf1\x62\x4f\xd5\x56\x38\xd2\xf5\xdb\xb0\x25\xc9\x64\xa4\xd7\x81\x57\x87\x9b\x72\x70\x75\x75\x84\x2a\xc9\x3e\x78\xba\xf3\x9e\x87\x6c\x8f\x17\x7e\xa8\xcf\xe8\x4b\xc8\x22\x81\x25\xd8\xf8\x35\xa2\xea\x81\x56\x47\x35\x4a\xad\x8f\x43\xc5\xb5\x5e\x82\x05\x8f\xf1\x38\x60\x1c\x2b\x4f\xe0\x3c\x7b\x9a\x6f\x34\xba\x7f\x64\x99\x22\x9f\x7c\x38\xd8\xbc\x07\x35\x9d\x1a\xf4\xfa\x4a\xde\x1c\x3c\xf2\xb3\xef\xa9\x1b\x22\x35\x7f\x02\x00\x00\xff\xff\x06\x22\x0c\x74\x91\x04\x00\x00")

func masterDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterDeploymentYaml,
		"master/deployment.yaml",
	)
}

func masterDeploymentYaml() (*asset, error) {
	bytes, err := masterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x4e\x03\x31\x0c\x40\x77\x7f\x85\x7f\x20\x87\xd8\x50\x36\x60\x60\x2f\x12\xbb\x9b\xb8\xd4\xf4\x62\x47\x8e\xd3\xa1\x5f\x8f\xaa\xa2\x5b\x90\x6e\xb5\xdf\xf3\x33\x75\xf9\x62\x1f\x62\x9a\xd1\x8f\x54\x16\x9a\x71\x36\x97\x1b\x85\x98\x2e\x97\x97\xb1\x88\x3d\x5d\x9f\xe1\x22\x5a\x33\xbe\xaf\x73\x04\xfb\xc1\x56\x7e\x13\xad\xa2\xdf\xd0\x38\xa8\x52\x50\x06\x44\xa5\xc6\x19\xc9\x2d\x59\x67\xa7\x30\x4f\x8d\xee\x02\xb8\xad\x7c\xe0\xd3\x1d\xa2\x2e\x1f\x6e\xb3\xef\x04\x01\xf1\x5f\x6f\x3b\x5f\x1e\xb3\x44\xb5\x89\xc2\x98\xc7\x1f\x2e\x31\x32\xa4\x3f\xe7\x93\xfd\x2a\x85\x5f\x4b\xb1\xa9\xb1\xfb\xd5\x63\x37\x3a\x15\xce\x68\x9d\x75\x9c\xe5\x14\x89\x6e\xd3\x79\x83\xe1\x37\x00\x00\xff\xff\x4f\x98\xa4\x7c\x24\x01\x00\x00")

func masterRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterRolebindingYaml,
		"master/rolebinding.yaml",
	)
}

func masterRolebindingYaml() (*asset, error) {
	bytes, err := masterRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8d\x41\xca\xc2\x40\x0c\x46\xf7\x73\x8a\x5c\x60\xa0\xff\xae\xcc\x29\x7e\x10\xdc\x87\xe9\xa7\x1d\xb4\x93\x90\xc4\x2e\x3c\xbd\xd4\x16\x5d\xb9\x0b\xef\x7b\xbc\xb0\xb6\x33\xcc\x9b\xf4\x42\xeb\x5f\xba\xb5\x3e\x15\x3a\xc1\xd6\x56\x91\x16\x04\x4f\x1c\x5c\x12\x51\xe7\x05\x85\xd8\x24\x8b\xc2\x38\xc4\xf2\xc2\x1e\xb0\x63\x73\xe5\x8a\x42\xa2\xe8\x3e\xb7\x4b\x64\x7e\x3e\x0c\x1f\x39\xb9\xa2\x6e\x1d\xc7\x1d\x35\xc4\xb6\x9b\x88\x55\x7f\x45\x55\x2c\x7c\xb7\xf2\xf1\x7d\x8e\xd0\x37\xd8\xd7\x42\xe3\x30\x0e\x07\x08\xb6\x2b\xe2\xff\x8b\x5f\x01\x00\x00\xff\xff\x10\x70\xf6\x36\xda\x00\x00\x00")

func masterServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceYaml,
		"master/service.yaml",
	)
}

func masterServiceYaml() (*asset, error) {
	bytes, err := masterServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8e\x02\x31\x0c\x05\xd0\x3e\xa7\xf0\x05\x52\x6c\xeb\x6e\xcf\x80\x44\xff\x95\xf9\x08\x0b\xc5\x8e\x1c\xcf\x14\x9c\x9e\x06\x51\xbf\x87\x65\x77\xe6\xb6\x70\x95\xeb\xaf\xbd\xcc\x0f\x95\x1b\xf3\xb2\xc1\xff\x31\xe2\xf4\x6a\x93\x85\x03\x05\x6d\x22\x8e\x49\x15\x64\xf4\x58\x4c\x54\x64\x9f\xd8\xc5\xfc\xda\x5e\x18\x54\x89\x45\xdf\x4f\x7b\x54\xc7\xfb\x4c\xfe\x72\xfb\x04\x00\x00\xff\xff\xe4\xf5\x04\x25\x70\x00\x00\x00")

func masterServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceaccountYaml,
		"master/serviceaccount.yaml",
	)
}

func masterServiceaccountYaml() (*asset, error) {
	bytes, err := masterServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\x0d\x02\x31\x0c\x05\xd0\x3e\x53\x58\xd7\x07\x44\x9b\x21\x28\xe9\xbf\x2e\x1f\x61\x41\xec\x28\x36\x14\x4c\x8f\xa8\xae\x7f\x98\x7a\xe3\x0a\x75\x6b\xf2\xb9\x94\xa7\x5a\x6f\x72\xc5\x60\x4c\xec\x2c\x83\x89\x8e\x44\x2b\x22\x86\xc1\x26\x3e\x69\xf1\xd0\x7b\x56\x7c\xdf\x8b\xd5\x27\x17\xd2\x57\x11\x81\x99\x27\x52\xdd\xe2\xef\xe5\xb0\x27\xf5\xb3\x79\x67\x0d\xbe\xb8\xa7\xaf\x26\xdb\x56\x7e\x01\x00\x00\xff\xff\xc1\xaf\xa6\x4c\x7c\x00\x00\x00")

func namespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_namespaceYaml,
		"namespace.yaml",
	)
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\x3d\x6f\xdb\x50\x0c\xdc\xf5\x2b\x88\xec\x8a\xed\x2d\xd0\x16\x34\x46\x96\x36\x28\x9a\xa6\x3b\xfd\x74\xb5\x1e\xfc\xbe\x4a\x52\x6e\x95\x5f\x5f\x08\xb6\x65\x09\x01\xbc\x84\x93\x70\x77\xbc\x23\x29\x89\x8b\xff\x05\x51\x9f\x53\x43\x5c\x8a\xae\x8e\x9b\xea\xe0\x53\xdb\xd0\x13\x4a\xc8\x43\x44\xb2\x2a\xc2\xb8\x65\xe3\xa6\x22\x0a\xbc\x43\xd0\xf1\x89\xc6\x86\x86\x58\x72\x9d\x0b\x84\x2d\x4b\xfd\x37\xcb\x01\x52\x11\x25\x8e\xb8\xc5\x69\x61\x87\x86\x72\x41\xd2\xce\xff\xb6\x9a\xdf\x7b\xc1\x24\xae\xb4\xc0\x8d\x21\x82\x12\xbc\x63\x6d\x68\x53\x11\x29\x02\x9c\x65\x39\xc5\x47\x36\xd7\x7d\x9d\xcd\x73\x73\x22\x35\x61\xc3\x7e\x38\x49\x25\x87\xe0\xd3\xfe\xad\xb4\x6c\xb8\x74\x47\xfe\xf7\xda\xcb\x1e\xa7\xb0\x33\xf2\x96\xf8\xc8\x3e\xf0\x2e\xa0\xa1\x75\x45\x64\x88\x25\x4c\x5d\xf3\xdb\x8c\x15\x16\xf3\xdc\x9c\x88\xe8\xb2\xe5\x58\x2e\x27\x63\x9f\x20\x53\x73\x4d\x2e\xc7\xc8\xa9\xbd\xba\xd5\xa3\xd5\xd5\x5b\xf6\x3a\xe7\xa6\xeb\x5d\xa1\x59\xd8\x58\x3e\xf2\xb8\xde\xf3\xf6\x65\xfb\xe3\xf1\xe7\xf6\x69\x22\x3e\xbe\xaf\x89\x12\x68\xee\xc5\x61\x16\x45\x14\x7c\xf4\xb6\x40\x88\x5c\xe9\x1b\xda\xac\xe3\x02\x8c\x88\x59\x86\x11\x5f\x7f\xf3\x33\x46\xf0\xa7\x87\x7e\xc2\x22\xf8\x23\x12\x54\xbf\x4b\xde\x61\xee\xd2\x99\x95\x67\xd8\xd2\xb8\xb0\x75\x0d\xad\x3a\x70\xb0\xee\x7d\x25\xe0\x76\x58\x0a\xb2\x58\x43\x0f\xeb\x87\xf5\x19\x4e\xb9\xc5\xeb\xe2\x7b\xbb\xa0\xb5\xe4\x80\xfb\x43\xbf\x83\x24\x18\xf4\xde\xe7\xd5\xe9\xce\x0d\xdd\xdd\x9d\xa5\x0a\x39\x7a\x87\x47\xe7\x72\x9f\xec\xe5\xc6\xef\xf0\x51\x7d\x4b\x59\xc4\x67\xf1\x36\x7c\x09\xac\x7a\xb2\xd5\x41\x0d\xb1\x76\xa1\x57\x83\xd4\x4e\xbc\x79\xc7\xa1\xfa\x1f\x00\x00\xff\xff\xc7\xb2\x22\x9e\xdc\x03\x00\x00")

func workerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerDeploymentYaml,
		"worker/deployment.yaml",
	)
}

func workerDeploymentYaml() (*asset, error) {
	bytes, err := workerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\xb1\x6e\x2c\x31\x08\x45\x7b\xbe\x82\x1f\xb0\x57\xaf\x7b\x72\x9b\x22\x7d\x14\xa5\x67\x3d\x24\x83\xc6\x63\x2c\xc0\xbb\x52\xbe\x3e\x9a\xd9\x6d\x53\xa5\xe2\x0a\x1d\x0e\x17\x52\x4a\x40\x43\x3e\xd8\x5c\xb4\x17\xb4\x2b\xd5\x4c\x33\x56\x35\xf9\xa6\x10\xed\x79\xfb\xef\x59\xf4\x72\xfb\x07\x9b\xf4\xa5\xe0\x4b\x9b\x1e\x6c\x6f\xda\x18\x76\x0e\x5a\x28\xa8\x00\x62\x35\x3e\x0f\xde\x65\x67\x0f\xda\x47\xc1\x3e\x5b\x03\xc4\x4e\x3b\x17\x24\xd3\xa4\x83\x8d\x42\x2d\xdd\xd5\x36\x36\xb0\xd9\xd8\x0b\x24\xa4\x21\xaf\xa6\x73\xf8\x61\x4a\x07\x9b\x75\x70\xf7\x55\x3e\x23\x8b\x02\xa2\xb1\xeb\xb4\xca\x4f\xa2\x3e\x5a\x38\x20\xde\xd8\xae\xcf\xed\x17\xc7\x39\x9b\xf8\x23\xdc\x29\xea\xfa\x17\xff\xc5\x83\x62\xfe\xf2\x66\x9c\xf6\x23\xcd\xb1\x50\x30\xfc\x04\x00\x00\xff\xff\x30\x78\x19\x41\x50\x01\x00\x00")

func workerRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRoleYaml,
		"worker/role.yaml",
	)
}

func workerRoleYaml() (*asset, error) {
	bytes, err := workerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\x31\x6e\xc3\x30\x0c\x45\x77\x9d\x82\x17\x90\x8b\x6e\x85\xb6\xb6\x43\x77\x17\xe8\x4e\xcb\x74\xcd\xda\x26\x05\x8a\x72\x01\x9f\x3e\x08\x12\x64\x09\xe0\xf9\xbf\xf7\x1f\x16\xfe\x21\xab\xac\x92\xc0\x06\xcc\x1d\x36\x9f\xd5\xf8\x40\x67\x95\x6e\x79\xab\x1d\xeb\xcb\xfe\x1a\x16\x96\x31\xc1\xe7\xda\xaa\x93\xf5\xba\xd2\x07\xcb\xc8\xf2\x1b\x36\x72\x1c\xd1\x31\x05\x00\xc1\x8d\x12\xa0\x69\xd4\x42\x86\xae\x16\xff\xd5\x16\xb2\x60\xba\x52\x4f\xd3\x15\xc2\xc2\x5f\xa6\xad\x9c\x04\x03\xc0\x53\xef\xf4\xbe\xb6\xe1\x8f\xb2\xd7\x14\xe2\xdd\xfc\x26\xdb\x39\xd3\x7b\xce\xda\xc4\x4f\xe5\xdb\x56\x0b\x66\x4a\xa0\x85\xa4\xce\x3c\x79\xc4\xa3\x19\x3d\xe0\x70\x09\x00\x00\xff\xff\x73\xce\x57\x9b\x2a\x01\x00\x00")

func workerRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRolebindingYaml,
		"worker/rolebinding.yaml",
	)
}

func workerRolebindingYaml() (*asset, error) {
	bytes, err := workerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8a\xc3\x40\x0c\x05\xd0\x7e\x4e\xa1\x0b\x4c\xb1\xad\xba\x3d\x43\x20\xfd\x67\xfc\x43\x84\xb1\x34\x68\x64\x07\x72\xfa\x34\x21\xf5\x7b\x98\x76\x67\x2e\x0b\x57\xb9\xfe\xda\x6e\xbe\xa9\xdc\x98\x97\x0d\xfe\x8f\x11\xa7\x57\x3b\x58\xd8\x50\xd0\x26\xe2\x38\xa8\x82\x8c\x1e\x93\x89\x8a\xec\xaf\xc8\x9d\xf9\xb5\x35\x31\xa8\x12\x93\xbe\x9e\xf6\xa8\x8e\xf7\x99\xfc\xe5\xf6\x09\x00\x00\xff\xff\xe3\x3c\x43\x66\x70\x00\x00\x00")

func workerServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerServiceaccountYaml,
		"worker/serviceaccount.yaml",
	)
}

func workerServiceaccountYaml() (*asset, error) {
	bytes, err := workerServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"aro.openshift.io_clusters.yaml": aroOpenshiftIo_clustersYaml,
	"master/deployment.yaml":         masterDeploymentYaml,
	"master/rolebinding.yaml":        masterRolebindingYaml,
	"master/service.yaml":            masterServiceYaml,
	"master/serviceaccount.yaml":     masterServiceaccountYaml,
	"namespace.yaml":                 namespaceYaml,
	"worker/deployment.yaml":         workerDeploymentYaml,
	"worker/role.yaml":               workerRoleYaml,
	"worker/rolebinding.yaml":        workerRolebindingYaml,
	"worker/serviceaccount.yaml":     workerServiceaccountYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"aro.openshift.io_clusters.yaml": {aroOpenshiftIo_clustersYaml, map[string]*bintree{}},
	"master": {nil, map[string]*bintree{
		"deployment.yaml":     {masterDeploymentYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {masterRolebindingYaml, map[string]*bintree{}},
		"service.yaml":        {masterServiceYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {masterServiceaccountYaml, map[string]*bintree{}},
	}},
	"namespace.yaml": {namespaceYaml, map[string]*bintree{}},
	"worker": {nil, map[string]*bintree{
		"deployment.yaml":     {workerDeploymentYaml, map[string]*bintree{}},
		"role.yaml":           {workerRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {workerRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {workerServiceaccountYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
