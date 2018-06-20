// Code generated by go-bindata. DO NOT EDIT.
// sources:
// deploy/data/virtlet-ds.yaml
package tools

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _deployDataVirtletDsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x5b\x73\x1a\x3b\xf2\x7f\xf7\xa7\xe8\x3a\xae\xfa\x27\x79\x18\x63\xa7\xfe\x67\x93\x43\xed\x3e\x38\x86\xe3\xa5\x62\x03\x05\xb6\x73\xde\x28\x8d\xa6\x19\xb4\x68\xa4\x59\x49\x33\x36\xfb\xe9\xb7\xa4\xb9\x30\x37\x30\x26\xb6\x6b\xc3\x8b\xb1\xa4\xfe\xa9\xef\xdd\x12\xf2\x3c\xef\x84\xc4\xec\x01\x95\x66\x52\xf4\x01\x9f\x0c\x0a\xfb\x55\xf7\xd2\x0b\x1f\x0d\xb9\x38\x59\x33\x11\xf4\x61\x40\x30\x92\x62\x8e\xe6\x24\x42\x43\x02\x62\x48\xff\x04\x40\x90\x08\xfb\x90\x32\x65\x38\x9a\xfc\x7f\x1d\x13\x8a\x7d\x58\x27\x3e\x7a\x7a\xa3\x0d\x46\x27\x3a\x46\x6a\x97\x1b\x8c\x62\x4e\x0c\xda\xef\x00\x55\x20\xfb\x69\x82\xd9\x0f\x27\x3e\x72\x5d\xac\x00\x50\x89\x30\xac\xb9\xac\x80\xb7\x9f\x95\xd4\x66\x8c\xe6\x51\xaa\x75\x1f\x8c\x4a\x30\x1f\x0f\x84\x9e\x4a\xce\xe8\xa6\x0f\x57\x3c\xd1\x06\xd5\x9f\x4c\x69\xf3\x83\x99\xd5\x3f\x33\x92\x7c\xe1\xa9\x83\x98\x8e\x06\xc0\xb4\x03\x00\x23\xe1\xe3\xc5\x27\x40\x41\x7c\x8e\xf0\x70\xab\xed\x88\x4e\x54\xca\x52\x2c\xf8\x00\x2a\x85\x21\x4c\xa0\x02\x85\xda\x10\xb5\x85\xfb\x68\x24\xf8\x08\x74\x85\x74\x8d\xc1\x27\x20\x22\x80\x8f\x9f\x3f\x59\x90\x1c\xd2\xac\x10\x12\x8d\x20\x97\x20\x34\x0a\x83\x0a\x98\x00\x26\x58\x05\xb6\x22\xde\x74\x34\xa8\x89\x76\x0a\xbe\x94\x46\x1b\x45\x62\x88\x95\xa4\x18\x24\x0a\x41\x20\x06\x8e\x53\xaa\x90\x18\x04\x62\xb1\x96\x2c\x8c\x48\x6c\xd1\x2b\xe6\xd9\x5a\x2d\x07\xd4\xa8\x52\x46\xf1\x92\x52\x99\x08\x33\xae\x99\xa5\xdc\x53\x0a\xbe\xb1\xe6\x80\x87\x5c\x03\xb1\x0c\x34\x48\xe1\xa4\x11\x32\x40\x0d\x8f\xcc\xac\xac\x47\x29\x32\xcb\xcc\xf6\x8f\x42\x5b\xce\xac\x39\x14\x59\x2e\xad\xa8\x9b\xad\x91\x2d\xf5\x65\x6b\x14\x40\xe1\xbf\x13\xa6\x30\x18\x24\x8a\x89\x70\x4e\x57\x18\x24\x9c\x89\x70\x14\x0a\x59\x0e\x0f\x9f\x90\x26\xc6\x3a\x73\x85\x32\xc3\x9c\x23\x47\x6a\xa4\xba\x43\x15\xe9\xfa\xb4\x07\x11\x31\x74\x35\x7c\x8a\x15\x6a\xe7\xff\xf5\x79\xbb\x62\x8d\x9b\x7e\x4d\x9c\xc6\x0a\x00\x19\xa3\x22\x46\xaa\x3e\x8c\x44\x6b\x32\x25\x3c\xc1\x16\xac\x05\x6e\xe8\xd6\xca\x7d\x55\xd8\xbd\x24\x38\x85\xbb\x15\x36\x9c\x02\xa8\x8c\x19\xea\x02\xe0\x83\x86\x25\xc7\xa7\x54\xf2\x24\x42\x08\x14\x4b\x4b\xbf\x39\xb5\x9e\x60\x2d\x13\xe0\x92\x24\xdc\x38\xfb\x3b\xab\xf1\x24\x64\x02\x02\xa6\x9c\x63\xa2\xd0\x89\x42\x0d\x66\x45\xb6\x1e\xec\xe8\x98\x72\xba\xb3\xdb\x59\xd7\xc2\x00\xfc\x0d\x70\xe6\xdb\xbd\xe1\xff\xca\x38\xc0\x27\xa6\x4d\xe1\x06\xd6\x5b\x4f\x0a\x29\xb3\xf0\x8e\x15\xc6\x44\xa1\x67\xed\x51\xaa\x82\x45\x24\xc4\x3e\x44\x4c\x11\x61\x98\xee\xd5\x73\x40\x3e\x3f\x4d\x38\x2f\x42\x78\xb4\x1c\x4b\x33\x55\x68\xa3\xa5\x5c\x45\x65\x14\x11\x11\x6c\x35\xec\x41\xaf\xba\xdd\x99\x5e\x95\x53\x99\x8e\x6e\xad\x7f\xeb\x2a\x41\xc6\xe4\xfa\xab\xf6\xb6\x9a\xf4\x32\x1d\x69\x2f\x60\xaa\x62\xbd\xc8\x12\x4f\x89\x59\xf5\xa1\x97\x6b\xd3\xab\x13\xb4\x70\x55\x22\x5a\x00\x4a\xc6\x24\x24\xce\x61\xe1\x1b\xcb\xd4\xcc\xa4\x20\x7c\xc7\x56\x55\x8c\x02\x37\x90\x74\x8d\x4a\x4b\xba\xde\x41\x94\x12\x65\x09\x7b\xd9\xc2\xb3\xda\xca\x02\x84\xcb\x70\x07\xb5\x35\x63\x75\xf6\x14\x96\x52\x65\xae\xc2\x44\xe8\x7c\x25\xdb\x82\x33\xbf\x97\xbb\x44\xcf\xd9\x4c\x67\xfe\xe0\xf2\x42\xcd\xe2\xc5\xa6\x29\x51\x1e\x67\xfe\x9e\x8d\xbd\xe6\x92\x52\x68\x4c\x77\x90\x55\x67\xbc\x96\x1e\x2c\x93\x4d\x07\xeb\x2e\x3e\x36\x13\xd2\x44\x31\xb3\xb1\xe1\x88\x4f\xa6\x1a\xbc\xb1\x62\x29\xe3\x18\x62\x50\x4b\xc6\x00\x28\xd2\xb6\x47\x7d\xbf\xff\x36\x5c\x8c\x27\x83\xe1\x62\x7c\x79\x3b\xac\xc0\xb8\xac\xf0\xa7\x92\x51\x3d\x31\x2c\x19\xf2\x60\x86\xcb\x66\xba\xa8\xd6\xea\xf4\xa2\x31\xe9\x88\x32\x49\x6d\x49\x3c\xb3\x1a\xb7\xd9\xbb\xc5\xcd\xc3\x68\x76\x77\x33\xbc\x5b\x0c\x46\xf3\xcb\x6f\x37\xc3\xc5\xf7\x87\xdb\xe7\x59\xca\xca\xc7\x2d\x89\xbf\xe3\xa6\x83\xb3\x9a\x02\xbd\x6c\x71\x63\x89\x4b\xa0\x01\xd3\xb6\xe8\x2d\xd6\x69\xd4\x98\x96\x71\xe6\xf8\x0d\x7d\x36\x99\x9e\xcf\x46\x93\x87\xc5\xfc\x7e\x3a\x9d\xcc\xee\xde\x8d\x6d\xad\x98\x4c\x17\x3a\x89\x63\xa9\xcc\x71\x8c\x0f\x26\x3f\xc6\x37\x93\xcb\xc1\x62\x3a\x9b\xdc\x4d\xae\x26\x37\xef\xa7\x73\xf9\x28\xb8\x24\xc1\x22\x56\xd2\x48\x2a\xf9\x71\x02\xdc\x4c\xae\x6f\x86\x0f\xc3\xf7\xe3\x9b\xcb\x90\x63\x8a\x47\xb2\x7b\x75\x79\x33\xba\x9a\x2c\xe6\xf7\xdf\xc6\xc3\xf7\x73\x14\x4a\x38\xa3\xd2\xd3\x89\x2f\xf0\x85\x8e\x32\xba\xbd\xbc\x1e\x2e\x66\xc3\xeb\xe1\x5f\xd3\xc5\xdd\xec\x72\x3c\xbf\xb9\xbc\x1b\x4d\xc6\xef\xc6\xbb\xcb\xd9\x0b\x85\x21\x3e\xc5\x0b\xa3\x88\xd0\xdc\x55\xa6\xe3\xf4\x3f\xbb\xfc\xb1\x18\x0c\x1f\x46\x57\xc3\xf9\xbb\x49\xa0\xc8\xe3\x22\x40\xdb\xbd\xea\x23\x83\x34\x4f\x89\x37\x93\xeb\xeb\xd1\xf8\xfa\xdd\xd3\x22\x97\x61\xc8\x44\x73\xc9\xa1\x1e\x3f\xbd\x5f\xdc\x4e\x06\xef\x18\xa1\x34\x4e\xbc\x48\x06\xc7\x86\x68\xe6\xf1\x15\x57\x9f\x2f\x06\xa3\x59\x93\xfb\x3e\xf4\xd0\xd0\xa2\x68\xe7\x9d\x45\xd1\x2d\xd3\x56\xa7\x5c\x36\x34\x59\x27\x72\x70\x97\x79\x0a\x23\x01\x94\x68\x84\x47\xdb\x68\xff\x0b\xa9\x01\x2e\x29\xe1\x65\x73\xeb\x10\xec\xec\x23\x11\xc6\x76\xd4\xf6\xd4\xc6\x0c\x08\x69\x40\x2e\x97\x8c\x32\xc2\xf9\x06\x48\x4a\x18\x77\x27\x3b\x29\xf0\x15\x9a\xd8\x5c\x90\x43\xfa\xd7\x6a\xb3\xa3\x37\xba\xb7\xd4\x3d\x1a\x2a\x99\xc4\xad\x56\xa7\x31\x5c\x27\xb5\x3d\x52\x24\x83\x84\xd7\xc2\x28\x23\x6c\x8f\x2b\x24\xc1\x44\xf0\x4d\xcb\xd8\x55\x48\x7b\x46\x6d\x61\x35\x06\x0f\x02\xaa\x37\xd1\x3f\xd3\x58\xff\x5c\x6f\xd8\x4d\xdd\x74\x3a\xd8\xe1\x8c\x6d\x6a\xdb\x9f\x3f\x43\xed\xd9\xc6\x1d\x8d\xae\xb8\xac\x3d\x66\x71\x19\xba\x03\x1c\x2b\x8f\x66\x2b\x54\x08\x3e\x52\xe2\xae\x15\xcc\x0a\xd5\x23\xd3\x58\x1e\xd7\x1e\x19\xe7\x10\x2b\x19\x24\x14\x01\x95\x92\xaa\x0a\xc9\xd9\x1a\xc1\xac\x58\xc5\xb1\x4e\xe1\x3e\xbf\xaa\x90\xf6\x04\xe7\xe5\x77\x0a\x74\x45\x54\x80\x29\x2c\x19\x47\xf8\x90\xe9\x40\x86\xbd\x34\xd2\x3d\xb2\x0c\xbe\xfc\xee\xfb\xbe\xf7\x15\xff\xf8\xe2\x5d\x5c\xe0\x17\xef\x8f\xdf\xff\x76\xe1\x9d\x7f\xfe\xff\xcf\xe7\x84\x9e\x9f\x9f\x9f\x7f\xee\x51\xa6\x94\xd4\x5e\x1a\x2d\xce\xcf\xb8\x0c\x3f\xf4\x61\x2c\x41\x27\x74\x95\x21\x4a\x55\x1e\x3b\x37\xed\x93\x43\xa4\xbd\xdd\x47\x96\x0a\x2b\xed\x83\x4e\xae\xcc\xe7\xa9\xdb\x46\x7b\xc9\xd1\xe3\x98\xc3\x83\x8d\x00\x26\x50\xeb\xa9\x92\x3e\x56\x49\xf0\x69\x7b\xc9\x95\x7d\x5a\xa9\x22\x63\xb1\xe7\x33\xd1\xab\xa4\x8a\x6c\xd4\xa3\x8d\x01\x2d\x29\x31\xe0\xc1\xfd\x78\xf4\x57\xbf\xe9\x80\xbd\xaa\xc3\x79\x4a\xc2\xdf\xad\x64\x3d\x91\x70\xde\x48\xb2\x9d\x47\xf5\xff\xf5\x24\x7b\x48\xf6\x7c\xbd\x34\x73\x9a\x25\x3f\x77\xbf\x52\xcd\xac\x40\x14\x96\x77\x5a\xe0\x6f\x40\x27\x31\xaa\x88\x89\x5f\x30\x29\xff\x8a\xc9\xb3\x8e\x92\x68\xc7\x83\x0d\x33\x77\x9d\xa3\x04\x1a\xd4\xe5\xcd\x4e\x7e\xa5\xd3\xcb\x5c\xa7\x67\x97\xb5\x36\x3a\xe0\xda\xa8\x5b\xee\x7c\x93\x5e\x2c\x83\xb6\x69\x2d\xaa\x9d\xe8\xbc\x7e\x3a\xa4\xc8\x1d\x9f\x2f\xab\x2b\x3a\x3a\xb0\x26\xa7\x6e\xd8\xb3\xdf\xbd\xca\xe1\xa1\x9d\x80\x9d\x34\xcf\xf3\x52\xd3\xc6\xaf\x94\x4c\xad\x33\xe6\xaa\x72\x57\x6d\x7b\xb2\x67\x45\xdd\x6f\x77\x09\x9a\x46\xfa\xa5\xbd\xe3\x7b\xc7\xf3\xcf\xd6\xf4\x43\x2f\x03\x77\x65\xcb\xfd\x79\x36\xd3\x58\xe5\x2a\xde\xa2\x56\xda\xad\xa5\x54\xee\xae\xdb\x1e\x3f\x21\x3b\x7e\x02\xa1\x14\xb5\x2e\xed\xed\x7e\xb8\xb1\xf8\x55\xc7\x6d\x73\xd8\x94\x66\x2f\x61\x77\x83\xdf\xd1\xde\xef\x45\xe9\x2a\x2b\x5d\x6a\xda\x0b\x52\xab\x19\xad\x32\xb2\x97\xb4\x5a\x54\x9b\x65\xf6\x14\xee\x26\x83\x49\x1f\x02\x29\x3e\x18\xb0\xdd\x26\x95\x01\xe6\xf7\xdd\x90\x25\x74\xd7\x3e\xd8\x44\xe3\xba\xde\x2d\xe1\x8a\xe9\xac\xd1\xcd\x4b\x2c\x5c\xcd\x46\xb6\xe9\x7d\xda\x00\x13\xda\x10\x9e\xa5\x27\xdb\x61\x54\x37\x64\x22\x33\xa5\xf3\x88\xed\x4f\x63\x67\x87\x88\xb2\xef\x9a\x7d\xc7\x4d\xfd\xb3\x78\x5d\x51\xd8\x15\x83\x07\x01\x35\x03\xaf\x2b\x1c\x9f\x07\xaa\x44\x68\xf3\xa7\x83\xbd\xc4\x3f\x51\x62\x0f\x2c\xb0\x07\x29\xa1\xb3\xda\xee\xac\xb5\x87\x40\x36\x0d\x53\xfb\xc5\xe2\x10\x7d\x96\x95\xb5\x9a\xdb\xba\x72\xe2\x41\x60\x7b\xad\xfc\x12\xb0\xae\xae\x6a\x5f\x4f\x75\x10\x77\x1d\x6a\x6f\x34\x04\x07\xf1\x25\xd0\x88\x26\x8e\x1b\xab\xd9\xae\xbc\xe6\xea\xef\x2a\x61\x5e\xd6\xb5\x74\x36\x2c\xfb\xdb\x9a\xe6\xdb\x09\xe5\x13\x7a\x46\x12\xb3\x92\x8a\xfd\xc7\xad\x39\x5b\x7f\xd5\x67\x4c\x36\x9e\x52\xe4\xef\x0f\x66\x92\xe3\x37\x26\x02\x26\xc2\x3d\x6f\x2a\x94\xe4\x98\xdf\xcf\x91\x98\x5d\xdb\xa4\xbe\x67\xa7\x13\x80\xd6\x1e\x2d\x48\x9d\xf8\xf6\xd8\xa5\xfb\x27\x5e\xbe\x7a\x5e\xfb\xc1\xff\xf0\x77\x1d\x56\x03\xed\xfd\x5e\xa6\x93\x23\x9e\x93\x28\x5b\x95\xec\x7a\xaf\xd4\x49\x5e\x9b\x3d\xf8\xed\x37\xf7\x45\xa1\x96\x89\xa2\x58\x8e\x97\x8f\x1f\x74\x3e\xe0\x9e\x28\xb8\xef\x29\x2a\x7f\xbb\xce\xdd\x6c\xe4\xff\x84\x68\x5e\xc3\xca\x1d\x32\x96\xec\x78\xb6\x49\x45\x55\xc8\xd4\x90\x28\x97\xa7\x26\x4d\x43\x96\x92\xfb\x8c\x5d\xfb\x97\x33\x9d\x7d\x79\x24\x86\xae\xde\x48\x82\x22\x7c\x12\x8d\xca\xce\xfc\xb4\x20\x9e\xed\xf1\x55\x96\x4c\x1a\x42\xbd\x69\xa4\x15\xe5\xc7\x3a\x84\xe7\xe7\xcb\x5e\x31\xec\x5a\xa6\xae\xc6\xdf\x4b\xc0\xaf\xf3\x8e\x2e\x83\xcd\x62\xa1\x9f\xb9\xf1\xdb\xa6\xa2\x68\x6b\xe4\x37\xd0\xcf\x2e\x47\xfa\x45\xd2\x94\x47\x55\xb0\xdb\xe9\x49\xcc\xb6\x4f\xea\x72\xcc\xae\x40\x48\xb4\x91\x51\x31\x18\xa0\x7b\xfb\x94\x97\xa2\x4a\x2c\xe4\xc9\xa9\xbd\x4d\x71\xda\x5c\x7f\xd5\x1d\xe8\xf9\xac\xab\x63\x11\x89\x63\x26\x42\x5d\x9d\x28\x3d\xb4\x98\xa9\x6c\x59\xe6\x92\x37\x8f\xc3\x9a\x3e\x5f\xdf\xbd\x2c\xec\xeb\xba\x54\xe3\x2d\x46\x27\xe0\x11\xd5\xed\xbf\x01\x00\x00\xff\xff\xab\xed\xf3\x6d\x93\x29\x00\x00")

func deployDataVirtletDsYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployDataVirtletDsYaml,
		"deploy/data/virtlet-ds.yaml",
	)
}

func deployDataVirtletDsYaml() (*asset, error) {
	bytes, err := deployDataVirtletDsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/data/virtlet-ds.yaml", size: 10643, mode: os.FileMode(420), modTime: time.Unix(1522279343, 0)}
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
	"deploy/data/virtlet-ds.yaml": deployDataVirtletDsYaml,
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
	"deploy": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"virtlet-ds.yaml": &bintree{deployDataVirtletDsYaml, map[string]*bintree{}},
		}},
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

