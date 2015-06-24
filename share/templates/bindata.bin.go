// Code generated by go-bindata.
// sources:
// index.html
// DO NOT EDIT!

// +build bin

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\x7b\x73\xdb\x38\x92\xff\x3f\x9f\x82\xc7\xd9\xbd\xda\x6c\xad\xa8\x4b\x6e\xf6\x2e\x95\xb5\x5d\x65\x5b\x72\xa2\x9a\xd8\xd6\xd9\x72\xee\xe6\xaf\x29\x4a\x84\x24\xc4\x14\xc9\x21\x28\x39\x1e\x95\xbf\xfb\x75\xe3\xc1\x27\x48\xf3\x65\xc7\x99\xca\x54\x4d\x99\x0f\xa0\xbb\xd1\x68\xfc\xfa\x21\x80\x39\xf8\xb7\xd1\xe5\xe9\xec\xd7\xe9\xd8\x58\x47\x1b\xf7\xd5\xd1\x01\xff\x63\x18\x70\x41\x6c\x07\x2e\xf0\x72\x43\x22\xdb\x58\xac\xed\x90\x91\xe8\xd0\xdc\x46\xcb\xc1\x3b\x33\xfd\x6a\x1d\x45\xc1\x80\xfc\xbe\xa5\xbb\x43\xf3\xff\x06\x37\xc7\x83\x53\x7f\x13\xd8\x11\x9d\xbb\xc4\x34\x16\xbe\x17\x11\x0f\xfa\x4d\xc6\x87\xc4\x59\x91\x4c\x4f\xcf\xde\x90\x43\x73\x47\xc9\x5d\xe0\x87\x51\xaa\xf1\x1d\x75\xa2\xf5\xa1\x43\x76\x74\x41\x06\xfc\xe6\x1f\x06\xf5\x68\x44\x6d\x77\xc0\x16\xb6\x4b\x0e\xdf\x28\x42\x11\x8d\x5c\xc2\xaf\xe1\x6e\xbf\xb7\x46\x76\x64\x5b\x1f\x7d\x16\x21\xf1\x87\x07\x03\xae\x80\xe4\xc1\x30\x69\x77\x74\xe0\x52\xef\xd6\x08\x89\x7b\x68\x52\x60\x69\x1a\xd1\x7d\x00\x72\xd0\x8d\xbd\x22\xc3\xc0\x5b\x99\xc6\x3a\x24\xcb\x43\x73\xb8\xb4\x77\xd8\xc0\xc2\x67\x85\xae\x2c\xba\x77\x09\x5b\x13\x12\x29\x02\x11\xf9\x1a\x0d\x17\x8c\xc5\xfd\xe1\x7a\x48\x3d\x87\x7c\xb5\xf0\xa9\xa4\xc0\x16\x21\x0d\xa2\x74\x97\x2f\xf6\xce\x16\x4f\xcd\xbc\x9e\x0d\x16\x2e\x80\xd0\x17\x36\x0c\x51\xc3\x21\x81\xab\xb7\xd6\x1b\xeb\xcd\x3b\xf5\xc0\xda\x50\xcf\xfa\x02\x3c\x1d\x18\xf8\x60\x63\x53\x4f\xb4\xdf\xef\xe9\xd2\xb0\x66\xc7\x1f\x3e\x8c\x47\x73\xea\x3d\x3c\x40\x3b\x29\x8c\xe8\xb1\xdf\x13\x97\x81\x86\x80\xc3\x70\x43\xdd\x5b\xf9\x92\xbf\xf0\x9c\x87\x07\x53\x29\xf5\x60\x28\x84\x93\xf2\x0f\xa5\x69\x1c\x1d\xcc\x7d\xe7\x5e\x3e\xf4\xec\x9d\xb1\x70\x6d\xc6\x0e\x4d\xb8\x9c\xdb\xa1\x21\xfe\x0c\x1c\xb2\xb4\xb7\x6e\xa4\x6e\x97\xf4\x2b\x71\x06\x91\x1f\x98\x46\xe8\xc3\x34\x62\x6b\xba\x02\x63\x81\x69\x88\xf9\x39\x34\x26\x86\x16\x01\x43\x22\xd0\xd3\xdd\x52\x47\xb5\xc9\xb5\x72\x42\x3f\x70\xfc\x3b\x4f\x71\x41\x09\x49\x98\x34\xe6\xc2\x6e\xa3\xc8\xf7\xb2\x42\x82\x20\xab\x95\x4b\xc0\xee\x5c\xd7\x0e\x18\x71\xd4\x4c\x8a\xc6\x52\xa7\xa2\x11\xca\x22\x5a\xa9\xc7\x76\xb8\xc2\x89\xfa\x49\xd2\x8a\x5f\xa7\xd8\xf2\xf9\x0e\xec\x98\x2d\x0b\x07\xbe\xe7\xde\x67\x9b\x40\xa3\x99\x90\x23\x51\x06\xe8\x1c\xba\x55\x50\x42\xbb\x1c\x00\xdb\x02\xa9\x17\xd3\x73\x28\x94\x98\x99\x04\xdb\xa0\x4e\x32\x5d\xf3\xd0\xf6\x40\xe5\xb9\x49\x54\x93\x22\xd5\x2a\x1b\x65\x66\x42\x35\x35\x0d\x3b\xa4\xf6\x80\x7c\x05\xee\x0e\x01\xca\x51\xb8\x25\xf1\xea\xcb\x4f\x04\x1a\x0c\xb2\x5f\x4b\x70\x28\x0c\xa3\x88\x1e\x07\x43\xe8\x94\x1f\x97\x9d\x19\xd2\xd6\x2d\x0c\x60\x43\xbc\xad\xb2\x6f\x71\xcd\xc5\x74\xed\x39\x71\x5d\xe2\xcc\xef\xf3\x2a\xc8\x71\x70\x69\x81\xa4\xb4\x68\x49\x34\x08\x09\x03\x50\xcb\x2c\x9b\xb8\x7b\xb8\xf5\x3c\xea\xad\x0e\x86\x2e\x2d\xd2\xad\xd1\x1f\x26\xa9\x30\x23\x80\x52\xf9\x76\x29\x7d\x8d\x26\xd7\xb3\xab\xc9\x09\xaa\xcb\x2e\x9a\x46\x41\x8c\xe1\xd6\xcd\x68\x30\xa3\xe3\x64\x9a\xf2\xcb\x2a\x81\x04\xf1\xc0\xa8\x58\x77\xe9\x69\x41\x68\x92\x4d\xe1\x8f\x46\xd7\xc8\x8c\x06\x66\x1e\x1a\x34\x63\x8e\x47\x3c\x99\xe2\x60\x75\x0a\xce\xa9\xae\xa8\x5d\x35\x3c\xd5\xc2\xb5\x75\x9c\x8d\xdc\x6a\x68\xb4\x00\x8a\x33\x15\x73\x05\x6e\x85\xb7\xa9\x71\x7d\x3a\xd6\xd9\xbc\x66\x9a\x8a\x6a\x6e\x63\xfd\x3a\x69\x3a\x9a\x3f\x27\xb1\x0d\x22\xba\x21\x85\x09\x8a\xc9\xd7\x24\x23\xe1\x4a\x50\xd3\xe0\x94\x76\x55\x64\xf4\x79\xc3\xbb\xea\x16\x86\xd4\xa9\x56\xc2\xec\x02\xd1\xb7\x2c\x2e\xa3\x32\x8b\x57\x97\x21\x5d\xad\xa3\x72\xa8\xa9\x32\xfb\x14\x6d\x97\xb2\x68\x40\x3d\x08\x83\x48\xe9\xaa\x52\xa4\x4b\x74\x2a\xe0\xf9\x27\x0c\x02\xf4\xda\xc3\x57\xcd\x34\xf6\x18\xaf\x80\x95\xb1\x0a\x58\xbf\x9c\x76\xab\x32\x4e\x3b\x7b\x05\x60\x1f\x3d\x85\x29\x14\x57\x8c\x61\x2f\x22\xba\x2b\xba\xb8\x9c\x07\xde\x91\x90\xa1\xf5\x97\xf8\xe0\x36\x6e\x17\xb3\x01\xf6\x7e\x38\xbc\xbb\xbb\xb3\xc0\x8f\x86\xf0\xbf\xb5\xf0\x37\x43\x11\x87\x43\xd4\xea\x12\x9b\x11\x36\x74\xed\x88\x30\x9d\x57\xb9\xbc\x9e\x8d\x2f\x66\x3a\x57\xd2\x0d\x6c\xd4\x58\x9f\x04\x71\x24\xf1\x3e\x20\xa7\x1f\x35\x72\x6a\x0a\x86\x3e\x8f\xaf\xae\x27\x97\x17\xcf\x83\x43\x39\x77\x9e\xba\x85\x1b\xc0\x0a\x99\x31\x28\x8f\xc4\x73\x01\x7d\xb4\x6f\x04\xb6\xe3\x40\x28\x83\xf9\x02\x23\x3b\x18\xf7\x3d\x5b\x6b\xd3\x84\xd0\xbf\xcb\xa7\x06\x48\x7b\x43\x36\x10\xa6\x8b\x3c\x54\x34\x9c\xbb\xfe\xe2\x16\xe3\xfd\xc1\xc6\x19\xbc\x55\x17\xfe\x72\x09\x39\xd7\xe0\x4d\x26\x84\xd8\xef\x23\xb2\x09\x50\xbf\x86\x09\x59\x0c\x88\xc5\xde\xbf\x17\x17\xbf\x81\xe1\x13\x17\xa8\x9b\x86\x95\x77\x98\x19\xb9\x32\xec\xfe\x33\x1b\xa1\xa4\x9a\x71\x72\x6f\x0a\xd8\x8c\x46\x9c\x8f\x7a\x84\x69\x1b\xbc\x07\x37\x52\xd0\x8f\x31\x8f\x20\x2e\x47\x56\xa6\xb1\xf4\x43\xb0\xe2\xed\x66\x73\x4f\x97\xa0\xce\x25\x5d\x2d\xd6\x64\x71\x3b\xf7\xbf\xf2\x77\x03\x86\xa9\xeb\x4f\xea\x5d\x45\xf0\x87\x34\xc5\x22\x10\xa4\x79\x3e\x49\x7e\x37\xcc\xa5\xed\x62\x34\xf6\xb7\x2f\x0c\xd2\x29\x61\x5f\xa7\x2e\x85\xc9\xb1\x3e\x52\x87\x08\xba\x93\xb3\xd7\x90\x7a\x0b\x0c\x52\xe9\xa4\x42\x48\xba\xac\x08\x25\x25\xa5\x99\x3d\x9f\x9c\x59\x33\x4c\xd9\xf5\x51\x25\xf5\x82\x6d\x24\xa0\x4c\x3f\xd6\xbc\xda\x92\x37\x22\xc3\x4b\xee\x39\x60\xac\xa9\xe3\x10\x2f\x46\x32\x7e\xa7\xb1\x7d\xd4\xfd\x2b\x9d\x24\x7d\xb3\xe3\xe6\x51\x1a\x3b\xc6\xd3\xa7\xf8\xc6\x93\x23\x08\x3e\x3e\x37\xb1\xa0\x82\x77\x2e\xe5\x4f\xf3\x04\xa3\xd9\x28\x36\x6b\x3f\xa4\x7f\xe0\x1a\x75\x07\xf8\x58\x8f\x60\x29\xbb\xc6\x46\x83\x55\xe8\x6f\x83\x01\x2e\x65\xe2\x94\xe0\x54\x66\xc9\x80\xdd\xf1\x2e\x46\x7c\x35\x60\x9b\x9c\x27\x12\xa9\x65\x89\x53\x2f\xac\x1d\x20\xc4\x89\xa9\x5a\x04\x8c\x99\x7a\x25\x7e\x3a\x3d\xa9\xb9\xa9\x2b\x6f\xff\x91\x2b\x51\x67\x1f\x75\x65\xaa\xb1\xba\xc6\xdc\xdd\x22\x9a\x89\xd5\xe5\x50\x86\x37\x4e\x9d\xf5\xd5\x7a\x64\xb9\x75\x29\x64\xc0\x18\x71\x72\xc6\x33\xa0\xca\x01\x6b\x13\x89\xca\x57\xad\xac\x67\xbf\xbf\xa3\xd1\x3a\xab\xac\x2b\xb2\x04\x57\xbb\x46\x29\xb5\x38\x1e\x8a\xf7\x1c\xbf\xa5\x02\xcb\x65\x1a\xa2\x24\x9a\xa5\x51\x88\xb8\xd1\xd9\x68\x8d\x32\x1d\x62\xd0\xe5\x80\x81\xc0\x8b\x75\x5c\x47\x9a\xe3\x9a\x8e\x53\xcb\xb3\xd9\xf1\xc9\xb5\x45\xcf\xa6\xc7\xa7\xbf\x8c\x67\xd7\xd6\x0d\xf5\x22\xdd\xd2\x14\x74\xed\x64\xda\x03\x7b\x71\x4b\xa2\xf2\x45\x31\x15\xef\xf5\x31\x40\x59\x14\xd0\x4e\xf6\xf1\xd5\xd5\xe5\x55\x03\xd1\x49\x18\xfa\x61\xb9\xe4\x63\xfe\xba\x17\xc1\x55\x60\x5c\x2d\xff\xc9\xaf\xb3\x71\x03\xf1\xe7\xf7\x10\x84\x95\x4a\x7f\x82\x6f\x9b\x0b\x5f\x8c\xbe\x4a\xb3\x70\x4d\x95\xaa\xda\x7d\x28\x53\x49\xa9\x07\x54\xc1\x41\x08\xf2\xba\x34\xf8\x68\xcd\xf1\xb5\x0e\x9f\x84\xcf\xe6\xaf\x01\x9d\x4a\x1c\x8c\x54\x3b\x3a\xf6\x0a\x93\x97\x8e\xbf\xe3\x02\x91\xc5\xcb\xcc\x88\x65\x38\xa8\xaf\x81\x94\xc7\x7b\x71\xff\xdf\x64\x38\xc9\xc3\xbe\x42\xd1\xb3\x62\x8a\x12\xdd\x4b\x5b\x6f\xa0\xfa\xd4\x6a\x7a\x2a\xcd\x4b\x16\xf5\x14\xff\xc8\xea\xce\xe8\x5d\x8c\xb6\xb5\xda\x45\xf7\x1e\xb4\x2e\x96\x68\x03\xa5\x27\x10\xf0\x54\x3a\x17\x1c\xea\xa9\xbc\x1a\x90\x32\x1a\xe7\x23\x6d\xad\x70\xde\xbb\xb5\xbe\x75\xe5\xf2\xec\x23\x6d\xb9\x57\x9b\x9c\x3d\x6f\xb6\xb4\x08\xb6\xe5\xe9\x52\xfc\xf2\xa9\xf2\xa5\xd3\xe9\x4d\x79\xc2\x04\xdc\x35\x53\x0e\x5d\x6a\x24\x46\x9a\x61\xfd\xa9\x33\xa3\x64\xa6\xda\xa4\x46\x62\x1a\x7a\xc9\x8d\x8c\xb9\x1f\x72\x5b\xf0\x21\x53\xd9\x7c\xd7\xe9\x52\x76\x31\xe5\xd6\xc5\xf7\x9b\x45\xc9\x35\x57\x96\x46\xf5\x3b\xb6\xd2\x3c\x0a\xa4\xf8\x1e\x12\x29\x2e\xe6\xd3\x65\x52\x7d\x44\xb5\x38\x61\xcd\x16\x7d\xd3\xe5\x1e\x3b\x59\x60\xd5\xd2\xc1\x42\xcf\x17\xe0\x5b\x9f\xb9\x12\xe9\x54\x54\x22\x9d\x27\xae\x44\x8e\x2a\x2a\x91\x4e\xbd\x4a\xe4\xa8\x49\x25\xb2\x38\xd6\x3f\xb5\xbf\x75\xba\x54\x22\x47\x3f\x2a\x91\xdf\x79\x25\x72\x54\x59\x89\xd4\xad\xaf\xd6\x23\x2b\xf5\xa0\xa3\xef\xa2\x12\x39\x7a\x81\x95\x48\xa7\xba\x9a\x37\x12\xc9\xa7\x73\x36\xb9\xb8\x1c\xd5\x2f\x87\x39\x30\x30\xdf\xa9\xa8\x87\x4d\xf8\xeb\x2e\xd5\x3c\xa7\x56\x35\x2f\x96\xbf\x51\x35\xcf\xf9\xde\xaa\x79\xb1\xbe\x53\xda\x29\x29\x6e\xe8\x66\xb4\xac\xba\x31\x6a\x57\xdd\xc8\xb3\xa8\x2c\x6f\xd4\xb5\xb0\x54\xe8\xa5\x46\xdb\x32\xfe\x52\xdd\x3b\x57\x94\x9c\x5c\x45\xa9\x86\xd2\x1f\xaf\x28\x75\xd5\x79\x8d\x8a\x52\xcd\x45\x91\xd1\x78\x97\x8a\x92\xf3\xfc\x15\xa5\xec\x8e\x03\x75\x53\x67\x5b\x71\x9d\x3d\x05\x01\xab\xda\x52\xf0\x2e\xb7\xa5\xe0\x6d\xb3\x2d\x05\x01\xd3\xee\x28\x78\xba\x11\x35\x1f\xc3\x93\x67\x0b\xbb\x55\x79\xb6\xa0\xde\x3d\x55\xb6\xf0\xf9\x43\x79\xb6\xa0\xdb\xd9\x75\xf4\x59\x6c\xea\x32\x56\xae\x3f\xc7\xb3\x09\x91\x1d\x6d\x35\xde\xa1\x90\x26\x14\x07\xf9\xa7\x4e\x13\xe2\x79\x6b\x93\x26\xf0\x49\x69\x97\x26\x88\x8d\x8a\x65\x05\x39\x03\xc3\x47\xdd\xbe\x48\x8d\xa9\xf7\x10\x02\x7e\xfe\xd0\x47\x08\xf8\xf2\x93\x9a\xf2\x7a\x61\xd9\xe6\x48\x41\xf4\x49\x52\x9d\x36\x91\xff\x53\x56\xa9\x76\x4d\xd7\x40\x43\xeb\x8f\xdd\xf6\x6e\xd5\xd2\x63\xef\x56\x2f\xc0\x57\x57\x1e\x92\x52\x1d\x76\x76\x68\xa0\xba\x8c\x43\x43\x86\x37\x0f\x0f\xff\xca\x9f\x56\x92\x87\x94\xf0\xdc\x12\x3f\xe2\xb6\xdf\x0f\xff\xfe\xea\xef\x43\x5c\x69\x62\xc0\x9a\x22\x9d\x4b\x03\x13\x06\x2f\x0c\x12\x48\x4f\x46\xc7\x51\x14\x82\x67\x80\x1b\x3e\x79\x0f\x0f\x1d\xa1\x13\xc4\xf8\xc2\xbe\x9e\xba\x3e\x03\xfe\x9c\x0f\x32\x14\xcb\x0a\x18\x9e\xf9\xa1\xe4\x98\x66\x88\x62\xe1\xdb\xf3\xff\xbd\xc6\xb3\x67\xa2\xc5\x2b\x91\x1b\xcf\x40\x47\x49\x1e\x7c\x24\x2d\xa4\xce\x60\x83\x2d\xc4\x05\xcc\xc4\x36\x1c\xba\xf0\x0e\xdc\xbe\x00\x23\x81\x48\xfc\x9c\xd8\x7e\x0f\x9e\x6e\x45\x8c\xbf\xd0\x7f\x18\x7f\x59\xf8\x21\x31\xde\x1f\x2a\x73\x9d\xde\x58\x9f\x28\x43\x01\xa2\x70\xbf\xbf\x25\xf7\x06\xaf\x91\x02\xa9\xf9\xfd\xe0\xc2\x14\xed\xad\x0b\x21\xed\x41\xe4\xa8\x25\x90\x00\xb0\xe1\xf9\x77\xa1\x2d\x77\x84\x83\xf8\x71\x8f\x83\x61\xe4\x94\xf6\x8a\xcf\xd4\xa1\xd5\x27\xaa\x52\xdd\x6f\x18\x09\x4f\xc5\xb3\x57\xb1\xe5\x27\xaf\x32\xf6\x7d\xd4\x85\xd1\xf5\x3d\x93\x7c\x8c\x3c\x23\x78\xd5\x1f\x9f\x89\xe3\x92\x92\x01\xe1\xab\x12\x46\xc3\x28\x4c\x59\x44\xf6\x4f\x1d\xfb\x10\x78\x80\x93\x8b\x17\xb1\xd0\x78\xf3\xc6\x60\x11\xac\x36\x92\x60\x0d\x8c\x2b\x75\x76\x14\x38\xab\xd5\x1d\xad\xd3\xeb\x3e\xbe\xc3\x17\x8f\x98\x03\x6f\x86\x13\x96\x39\x8d\xb6\xf5\x68\xc6\x6d\x1f\xfd\x35\x87\x56\x6d\xb8\xc0\x6c\x3d\x3d\x13\x9c\xaa\x2e\x5c\x86\x52\xab\xfc\xa9\x3a\x8c\x19\x25\xa7\x31\x1f\xf9\x01\x82\x2f\x77\x81\xee\x51\x82\x8e\x7c\x02\x9b\x00\x87\xca\xec\x1a\xa2\x87\x43\xd9\x6d\x82\x1c\xa3\x33\x4e\xa4\x80\x1e\x2a\xe9\x14\x08\xe2\xd0\x90\x1f\xd8\x13\xbd\xad\x11\x0d\x2f\xf8\xf1\xbc\xdc\x32\xca\x41\x48\x85\x12\x38\x92\xfe\x0d\xff\x18\x6f\xde\x22\xb7\x81\x60\xe7\x90\x5d\x86\x11\xd9\x21\xa3\xd7\x5a\x14\x6a\xc5\xed\xbf\xd2\xcc\xb4\xa3\xd2\x33\x2b\x20\x04\x2c\x7b\xde\xeb\x78\x67\x53\xb7\x59\x17\x58\x49\x30\xc3\xff\xee\xcd\x59\xf0\xaf\x03\xb6\x0d\x12\x24\x29\x1f\xc3\x96\x91\x80\x84\x0b\x08\x4a\x60\x24\xc9\x8d\x11\x53\x9c\x8a\x07\x5c\x7a\x45\xb3\x16\xd4\x29\xa9\x66\x7e\x64\xa7\x07\xd2\x09\xb4\xb2\x55\x87\x8e\xc8\xa5\xd5\x4b\xb4\xc6\xa3\x22\x78\x1a\xdc\x34\xd4\x95\xb4\xe9\x4f\x70\xc9\xac\xa9\x1d\xda\x1b\x66\x8d\x2f\x6e\xce\x2d\x67\x69\x98\x67\xd7\xa6\x61\x9a\x25\x19\x40\x63\x6a\xe7\xd3\x3e\xa9\x1d\x7f\x3e\x9e\x7c\xea\x8d\xda\xcd\xf5\x78\xd4\x1b\xb1\xd9\xe5\xec\x18\x45\xeb\x04\x7a\x19\xa8\xea\x03\xf9\x64\x11\xb1\x23\xf4\x09\x2a\x1a\xec\x93\x25\xce\x67\x03\x3f\xc9\xef\x99\xd0\xaf\x84\x5b\x2b\xf8\x9b\x40\xe6\x4c\x1a\x76\xd9\xf6\x8e\x7f\x48\xb2\x33\x00\x8a\x9f\x45\xfa\x43\xc0\x4c\xa1\xbb\x7b\xf0\x16\xff\xf2\x9a\xfb\xb6\xc2\xd1\x88\x7f\xa5\xa3\x24\x0c\x2a\xb4\x3e\xf7\xb7\x5e\x44\x9c\xca\xe6\xda\xba\xcc\x11\xf7\x74\x2d\xfa\xa1\xbb\x6b\xd1\x8d\xfb\xa3\xde\xc2\xae\x2c\x64\x74\x87\x20\xda\x2e\xf8\xa2\xcb\x04\x7f\x26\x25\xa1\x17\xcd\x84\x5e\x72\x89\xd2\xa5\x95\x41\x9d\xe6\x30\x43\x55\xd8\x93\xa5\x58\x7f\xbd\x43\x87\x11\x71\x23\x7b\xe2\x35\xee\x72\xb9\x8d\x9a\xf4\x69\xc6\x21\x4b\xbc\xd3\xa2\xa5\x7d\x86\x2d\x99\x84\x6b\x02\x8b\x2e\x5c\xda\xa5\xeb\xb4\x98\xae\x18\xf2\x47\xa6\x93\xc9\xec\xda\x00\xb8\x33\x18\x96\x63\xd3\x5f\x60\x99\x78\xd5\x49\xcc\x41\xee\xa4\xe8\xfc\x60\x98\x7e\x72\x84\x47\xaf\x5b\xe6\x52\x8f\x0b\x07\xb3\xf2\xcd\xa5\x8b\x10\x42\x0c\xfe\x43\x9b\xb1\xf1\x9d\xad\xeb\x1b\x3f\x7f\xe8\xa0\xc0\x93\x9c\x88\x7f\xfd\xf9\xc3\x93\xcb\xd8\x58\x8f\x4d\x85\xec\x82\xaa\xb4\xdf\xb8\x2e\x3e\x6e\xd0\x09\x55\x05\x0d\x0d\xac\xca\xa3\x10\xfd\xe3\xaa\x24\xfc\x03\x58\x1f\x9f\xd9\x97\x84\xac\x25\xa0\xaa\x02\xd3\xaa\x45\xd7\x03\x38\x95\xa1\xe6\x33\xb1\xd7\xe3\x61\x1d\xe6\x7d\xa0\x5e\x09\xd2\xf5\xc1\xbe\x1b\x9e\xa5\x01\xa8\x0f\x40\x53\xc7\xd6\x3a\x21\x9a\x24\xa2\x81\x34\x75\xaa\xae\x7f\x4c\x53\x94\x7f\x80\x5a\x8d\xd9\xfd\x81\x6a\x2f\x04\xd5\x44\x34\xf5\x6d\xb1\xad\x5c\x86\x17\x81\x70\x19\x44\xea\x0e\x71\x1b\xb2\x69\xf1\xf3\x25\xf4\x4a\x00\xee\x7c\x7c\x5e\x80\x36\xfc\x72\x8c\x00\xb5\x5b\x8a\x9f\x1f\xc4\x1e\xd6\x2f\x14\x85\xca\xa1\x5a\xea\x4d\x4d\x14\xc0\x0e\x67\x4d\xca\x56\xd8\xa1\xd7\xa2\xbd\x24\xd8\xa5\x64\x85\x24\x7a\x2d\xd9\x03\xc1\x6f\xf0\x3b\x63\x6a\x75\xe0\x94\x34\x69\x5f\x51\x57\x7a\x86\x82\x92\xb2\xfb\xee\x2b\x28\xfe\x96\x11\x10\xd2\xee\xed\xcb\x6d\xf7\x29\xdb\xc2\x27\x87\x62\x1b\xb8\xaf\xe5\x64\xca\x7f\xa6\xc6\x2d\x0a\xba\xf2\xfe\xc9\xe5\xe5\x27\x8b\xad\xfd\x3b\xb1\xc3\x0b\x97\x63\x61\xe3\x5e\xbc\xb7\xbb\xf0\xc6\xe4\xfb\x22\xc4\x66\xa5\x8f\x21\x59\xd6\xe7\x92\xfc\x6c\x7e\x4e\x36\x7e\x78\xaf\x76\xea\xa5\xb6\xca\x25\xdb\x77\x32\xdf\x91\x52\xfb\xd7\x5a\x0e\x0e\x46\x91\xdb\xc3\x63\x26\xb2\x88\x2d\x6b\xa9\x9d\x1e\x35\xce\x92\x42\x73\xfc\x7b\xbc\xc0\x4f\x89\x95\x49\xf2\x3f\x5b\x12\xde\xc7\x7c\x78\xb2\xae\xdd\xfc\x27\x25\x92\xfb\x57\xd8\x76\xbe\x49\xbb\x84\xfa\x7b\xce\x6a\xed\x34\x13\x5f\x42\xe3\xf7\xf9\x9d\xa3\xf5\x6d\x07\x24\x26\x5c\xb1\xb1\x91\x08\x3b\x49\x6f\x44\x6b\x64\x26\x92\x60\x4a\x59\x52\x28\xb5\xcd\xcc\xee\x24\x2b\xbb\xb3\x83\x7e\x85\x45\x8a\xbc\x0f\x5d\x8e\xe4\xf9\x9b\xd6\x03\x24\x06\x52\x2b\x8e\xb1\xc1\xf9\xc2\x6a\xbb\xa8\x42\xb2\x64\xf7\x63\x51\xf2\xe9\xf8\x6a\x72\x39\xb2\x64\x13\x2e\x7e\x61\x07\x59\x6a\xb3\x5e\x66\x0b\x99\xbc\xcc\xaf\xe7\xc6\x16\x56\xbe\x6e\x1f\xc1\xe7\xf4\x36\x3a\x8d\x58\xfc\xa2\x29\x4e\x07\xec\x1b\xc2\x74\xc0\x9e\x01\xa5\x83\xd4\xde\xa6\x69\xe8\x2f\x08\x44\x73\xac\x26\x4e\x07\xac\x3d\x4c\xe3\xe0\x9a\xa1\x74\xfd\xcd\xc5\x2f\x13\xac\xfb\x59\x94\x01\xab\xfc\x02\xe2\x0b\xf2\x17\x38\xc3\x7d\x22\x70\xca\x52\xd3\x60\x5a\xe5\x2d\xf2\x82\x4e\xaf\x39\x48\xc0\x5f\xcf\x8f\x46\x64\x11\x12\x9b\x43\x39\x4c\x88\xc4\xf4\x0a\x51\x5d\x58\x1a\x55\x82\x9e\x5e\xde\x5c\xcc\xac\x80\x79\x45\x31\x07\xed\x24\x4c\x0e\x7e\xd6\x12\x10\x22\xac\x4a\x4d\x56\x08\xa8\x8e\x2b\x25\xec\x03\x77\xcb\xd4\x26\xd7\xa6\x7e\xea\x29\x7d\x04\x5a\x55\x3b\x17\x11\xb0\xbe\x3d\x04\x6b\x98\x0a\x07\x80\xaf\x49\x1e\xac\x74\x9d\xcf\x85\x03\x55\xdf\x0b\x28\x66\xc2\xd8\xc9\x9a\x4e\x46\x98\x90\x79\xab\x8a\x2d\xbd\xd2\xd5\x80\x36\xe3\x3e\xf5\x52\xde\xb8\xc7\x4d\xae\x87\x7c\x9d\xfa\xaf\x49\x55\x31\x60\x90\x03\x87\xda\x51\xc8\x7b\xdc\x63\x5a\xb3\xc4\x98\x8c\x2a\xa4\x10\xad\x47\xf7\x0d\xbb\x5d\xd0\x45\xcd\x02\x40\xdc\xe5\x9a\xfe\xd1\xb4\xcb\x15\x61\x60\xa2\x5e\x45\x3d\x13\x0b\x00\x6a\x83\xc6\x91\xea\x36\x93\x1f\x2d\xaf\xde\xf5\xd3\x68\x02\x7e\xe6\x13\x80\xd5\xdc\xaa\x09\xc8\xd7\x78\x3b\x55\x14\xd4\xfa\xd2\x16\x14\xde\x3e\xe3\xf6\x3f\x84\x08\x18\x69\x2f\x7b\xe2\x90\xd6\x4d\x9f\xb4\xae\xc7\x57\x7d\x6d\x25\xe4\xe3\xbc\x9a\xf4\x46\xeb\x62\x72\x3a\xee\x8d\xd8\xe7\xc9\xd5\xac\x37\x62\x57\xe3\xeb\xde\x68\xcd\x26\xe7\x63\x0c\x3c\xc5\x42\xec\x4d\x77\xc7\x9c\x6a\xd7\x9d\x93\xd2\xa5\x74\xaf\x32\xa9\x80\x52\x66\x2f\xe0\x82\xe4\x87\x85\x70\x4d\x1a\x96\x3c\x04\x23\x03\xc0\xd4\x35\x84\x80\x86\x6c\x2a\x1e\x98\xd2\xe3\x64\xfe\x61\x9a\x54\x7b\x08\x32\xd5\xb7\xd5\x8f\xe4\xc1\xbb\xb8\x9a\x7d\x14\x1f\xe9\x89\x83\x78\x1e\x9f\xe2\xe1\x59\x08\x33\x63\x46\xe2\x40\x4e\xea\xf4\x91\x69\x80\x7e\x16\x64\xed\xbb\x10\xe5\x1f\x82\x33\x9d\x26\xb7\x3c\xd4\x09\x08\xb8\x01\x07\xe1\x4b\x04\x0c\xa9\x87\x9f\x6d\x77\x1b\x3f\x7d\xc5\xf5\xad\x39\xfa\xd3\xd4\xd7\xa7\x6a\xbb\xd0\x5b\x1e\x17\xb1\xe2\x90\x25\x3e\x17\xc4\xb9\x3f\x3c\xa8\x03\x05\x0d\x38\xec\x56\xfc\xc7\x48\x0e\x9f\xa1\x01\x91\x00\x9e\x29\x13\x8f\x62\xaf\xe0\xbb\x48\xf5\xd0\xfc\x67\xec\x40\xe4\x07\xe7\xc5\xb1\x58\xfe\xdd\xd6\x52\x38\xaf\x21\x41\x1c\xcf\xc8\x80\x65\x63\x2f\xd6\x49\xc0\x22\xb9\x9c\xc3\x43\xec\x57\x08\x5c\xe4\xbf\xb8\x20\xa3\x97\xed\x96\xfb\x1d\x24\x61\xdd\xf0\xc8\x42\x17\x55\xd4\xf5\x67\xff\x9d\x90\xcf\x13\x7e\x5d\x1e\xb3\x60\xe1\x1c\x9b\x5d\xd8\x39\x07\xab\x6d\x05\xe9\xef\x8e\x3a\x24\xac\x6a\x59\x4f\xda\x77\x89\xb4\x78\x44\x99\xa4\xc5\x95\x97\xd7\xf8\xbc\x52\xf4\x7a\xac\xfe\xf9\x1f\x09\xaf\x25\x75\x49\x60\x47\x6b\x0d\x3b\x39\x77\xd8\xe2\x37\x6c\xf2\xa8\xd7\xaf\x61\x2d\xf2\x63\x26\x35\xe2\xdf\xd2\x83\x97\x69\xc3\x75\x5e\x97\x80\x71\x6e\x81\x18\x19\xd2\x55\xed\x63\x2c\x6d\x1b\xd2\xa8\xb3\x99\x71\x48\x23\x4e\x7a\x66\x3f\x48\xd0\xe5\x07\x13\x9a\xff\x49\x23\x79\x85\xb1\x5b\xe9\xcb\x40\x9a\x6a\x69\x03\x6e\x76\xa5\x6f\x1d\x1a\x92\x45\xc4\x0b\xf2\x3d\xfd\x3e\x12\x5b\x43\x07\xd7\x95\xf3\x59\x45\x67\x53\xed\xa8\x3a\xbb\xa9\x74\x09\xa5\xa1\x8b\xe2\xdf\xfb\x10\xb9\xb8\xb4\xf7\xd4\x57\x49\xc4\xf3\xf8\x23\x24\x53\xee\xa6\x5e\xe3\xa7\x0e\x76\xe8\x2d\x78\x67\xf1\x30\xb6\x4d\xb3\x59\x42\x9a\x0a\x57\xd0\x54\xd7\x99\x1f\x37\xc4\xa6\x6b\xe0\x71\xec\xd2\x95\x27\x8f\x33\x0a\x06\xf2\x2b\x3c\xf0\x0e\xeb\x06\x99\xe3\xaf\xca\xbb\xc9\xdc\x3a\x75\xe6\xb5\x70\x56\xd2\x3a\xb5\x43\x12\xe5\xce\x49\x66\x4f\x46\xda\xd2\x16\xd6\xf1\x70\xfe\x3f\x00\x00\xff\xff\x75\xf0\xb9\xed\x0b\x73\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 29451, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"index.html": indexHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

