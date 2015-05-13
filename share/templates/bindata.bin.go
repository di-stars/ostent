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

func bindata_read(data []byte, name string) ([]byte, error) {
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

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _defines_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5f\x6f\xdb\x36\x10\x7f\xef\xa7\x20\x84\x74\x58\x81\x46\xc1\x86\xee\xa5\x73\x04\xb8\xb1\xd3\x0a\x8d\x93\xc0\x7f\x02\xec\xa9\x50\x2c\x3a\xe6\x22\x8b\x1a\x49\xb9\xf3\x02\x7f\xf7\xf1\x28\x51\x7f\x2c\xd1\x96\x64\x67\xd8\x43\xf3\x50\x8b\xe4\x4f\xc7\xe3\xdd\xfd\x8e\x47\xaa\x2f\x2f\x3e\x5e\x90\x10\x23\x2b\xf9\xe5\x1f\x3f\x26\x0f\xdf\xe6\x51\xcc\xe8\x77\x6e\x6d\xb7\x2f\x2f\xdf\x89\x58\x22\xfb\xee\x61\x38\x1e\xbb\x83\x21\xf4\xd8\xf0\x0f\x0e\x38\x86\xdf\xb3\x2b\xf4\xf1\x12\x51\x86\xec\xab\x9b\xfe\x64\x72\xdb\x1f\x0d\x91\x35\x0f\x3c\x9e\xbc\xcd\xbc\xf0\x09\xa3\x33\xf2\x1e\x9d\xcd\x29\xc3\x80\xb5\x07\x9e\xf0\xec\xab\xfb\x99\x7d\x43\xb8\xd8\x6e\x7b\x82\xa1\x67\xbc\xb9\x94\xb2\x00\x62\xdf\x6e\xb7\x4e\x4f\xf8\x48\x09\xb9\xb4\x04\xfe\x5b\x9c\x33\xf2\xb4\x14\x28\xa4\xdf\x99\x17\x59\x4e\x01\xd9\xbb\x10\x7e\x3d\xda\x72\x7a\x3c\xf2\x42\x04\x2a\x6e\xb7\x99\xf4\x19\xc7\xec\x0a\xb0\x72\x96\x62\x1f\x88\x02\xbc\xd3\x5a\xe2\x64\xc3\x53\x81\xc8\x29\xf4\x75\x17\xe8\xfa\x01\xde\x55\x11\xfa\x76\x24\x5e\x08\x26\x87\x71\xe8\x2b\x77\x14\x7e\xde\xec\x75\xac\xf0\x1e\x03\x6c\x81\xd9\xe1\x21\x53\x0a\x1a\xbf\x20\x2e\x18\x89\xf0\x79\x82\x91\x4a\x2f\xb1\x07\xba\x33\x78\x84\x19\x97\xf0\xb0\xc7\x33\x60\xc9\x64\x4d\x29\x26\x0e\x89\x5c\xe6\xdb\x5c\xf1\x43\x12\xa4\xe5\x8e\x13\x00\x96\x3a\x28\x01\x6c\x07\x4f\xc9\xf2\x1e\xa9\xbf\x91\xa6\x14\x78\x15\x05\x9e\xd8\xc3\x07\x64\xab\x88\x53\x78\xf9\x0b\x56\x72\x0e\x1b\xdd\x5f\x3c\x6e\x04\xe6\xdf\x4e\x4b\x29\x9f\xf0\xe7\x9c\x4e\x83\x6b\x35\x47\x95\x52\x00\xb3\x07\x84\xdd\x7a\x2b\xfc\x15\x6f\xca\xdc\x2a\x10\x2a\xc5\xf1\x67\x00\x7e\x99\x8e\x6e\xaa\xdc\xaa\xa2\x99\x19\x5c\x8c\x72\xfd\x42\x7f\xed\x91\xa0\x19\x54\x06\x92\x34\xea\x4f\xe1\x23\x8f\x7e\xef\xf1\x38\xaa\x30\x45\xa3\xee\x31\x9b\xe3\x50\x14\xf8\xb2\x33\xb2\xdd\xe6\xae\x57\x82\x1a\xcc\x3e\xa5\xc2\xcb\x14\xed\xc2\x33\xed\x72\x4d\xb6\x03\x4e\x6d\xc7\xc5\x34\x80\xce\x6e\x72\xf7\xdf\x90\xf0\x19\x32\x4e\xce\x0c\x40\x63\x66\x39\x28\xff\xeb\x79\x68\xc9\xf0\x02\xac\x77\x23\x03\xe6\x7a\x62\x7f\x91\x4d\x99\xb6\xe4\x5f\x6e\xd8\x74\x48\xa7\x34\x84\x9c\x01\x5e\x93\x79\xca\xa9\x1a\xa0\xc7\xb0\x28\xa0\x33\x63\x7b\x15\xb6\x1e\xd6\x69\x74\x6f\xd4\x49\x0e\x15\x66\x19\xd1\x38\x14\xd8\xdf\x8d\x0a\x0d\x6c\xab\x13\x2a\x25\xe5\xb2\x4e\xfd\x87\xbe\x7b\xa3\xd5\x2a\x4f\x95\x0c\xe9\xd8\x53\xe1\x9d\x9a\xa9\x1e\x58\x50\xeb\x38\x9d\x66\x93\xe1\x20\xb3\x54\x79\x2a\x35\x94\x6d\x48\xc0\x23\xad\x52\x2d\xb0\x68\xa9\xe3\x74\x9a\xde\x4d\xfb\x06\x3b\x25\x43\x7a\xe5\x8a\x5d\xf5\x76\x4a\x81\xfb\xec\x94\x92\xaf\x6d\x0a\x2f\x25\xe1\xae\x79\x9c\x84\xd4\x7f\xf5\x44\x9e\x4c\xf2\xff\xce\xe4\xee\x82\x61\xdc\x10\x1a\x37\x4a\xe5\x00\x33\xe4\xf2\xc2\x50\x97\x64\xee\x2a\x7b\x1e\x97\xcd\x53\xc7\x77\xae\x9d\x76\x33\x60\x9a\x52\x0d\xf9\x51\x27\xb7\xc3\x14\x4c\x52\xce\x61\x9c\x4a\x03\x87\x61\x09\x33\xbb\x54\x48\x65\x6e\x74\xe3\x17\x69\x59\x27\xe9\xb3\xc5\x22\x67\x8f\x6b\x2a\x83\xc8\xc2\x2e\xf1\xc6\xc9\xfb\x1a\x45\xbd\xc4\x0e\x70\x20\x3c\x37\x6c\x0c\xbd\x8b\x45\x13\x6c\x33\x89\x05\x61\x5d\x22\x98\xec\xd4\x23\xed\x8b\x7f\x57\x46\x24\x5b\x78\xd5\xa0\xad\x96\xe0\x48\x10\x11\xe0\x4b\xeb\x93\x3b\x9d\xa0\x48\x86\x18\xc7\x73\x1a\xfa\xb2\x34\x0f\x6b\x0a\xf3\x1e\x71\x1e\x7b\x17\xc4\x89\x78\xe3\x1a\xdf\x38\x81\xb4\xd2\x69\x67\x10\x40\x08\xf4\xe9\x8f\xe9\x70\x82\x56\xd4\x8f\x03\x8a\x3e\x7c\x36\x2f\xe4\x13\x4c\xf3\xf6\xc3\xe7\x53\xcd\x63\x5c\x4f\xed\x44\x2d\x29\x4b\x8e\xdf\x11\xc9\x02\x33\x46\xd9\x91\x94\x4d\x64\xfc\xe0\xac\xc9\xb8\xff\x29\x69\xcb\x7c\xd5\xbb\x76\x25\x06\x3b\x70\x69\x87\xa8\xa7\x14\x5d\xa2\xa6\x51\x70\x17\x62\x96\xc9\xd8\x58\x74\x6b\x2a\x16\x79\xd4\x95\x8b\x91\x37\x7f\xc6\xe2\x48\x32\xa6\x42\x7e\xb0\xd1\x68\xde\x1f\x74\x6c\xb6\x9b\xbd\x1e\x29\x2b\xf2\x5f\x95\x9a\x25\x5a\x75\xe3\xe6\x0a\xaf\xda\x92\xf2\x3d\x3a\x93\x6f\xe5\xcc\x1c\x0d\x47\x55\x4e\x4a\x84\xfd\x95\xc0\xe4\x9a\x91\x79\xcf\x01\x46\x00\xf0\xba\xc9\x21\x0e\x80\x95\xdb\xb8\xbc\x3f\x3d\x98\xe9\x14\xd0\xe0\x50\x06\x2f\x1e\x7d\xc1\x26\x85\xbc\xc6\x45\xb6\xe5\x80\x4d\xf6\x8d\xd7\x9d\xa4\x4e\x73\x84\xd2\x41\xd2\x2d\xc4\x22\xde\x32\xc2\x22\x46\xe7\x79\x78\xdd\x4f\xd4\x44\xd5\x10\x03\x98\x7d\xef\x0e\x4c\x1f\x65\x2c\x07\x95\x50\xfb\x9c\x9f\x21\x67\x45\x64\xf1\x3a\x30\x47\x70\xcc\x0e\x6f\x2b\xf9\xd4\x8c\x50\x46\xc4\xa6\x21\xfc\x56\x1e\xbe\x1b\x42\x27\xe4\x9f\xa6\xd0\x31\xe6\xc4\x57\x97\x14\xb5\x70\xa0\x09\x9c\xee\x35\x7c\x4a\x56\x35\x92\xf5\xd5\x4c\xc9\x1a\xbb\x9b\x6c\x17\xc6\x44\xbc\xfb\x65\xf4\xaf\x27\xb9\x8c\x2e\x5d\x35\x94\x6f\x7e\xef\x27\x32\x78\x6a\x2f\x34\x93\x91\xec\x86\x52\x36\xb2\xeb\xcc\x3a\xdc\x11\xd7\x99\xbb\x1a\xcd\x8c\x1a\xcd\x8a\x1a\xcd\xcc\x1a\xcd\xda\x6b\x54\xba\x1c\xaf\x68\x34\x19\x8e\xeb\x2e\x58\xd3\x11\x7d\x6f\x06\x8d\xda\x3b\x5f\x8d\xeb\x7e\x0d\x5d\xf1\xda\xd8\x35\x79\x4d\x8e\xe4\x5e\xcb\xf5\xa9\xc5\x9d\xd0\x6b\xb7\xee\xd5\xb0\xde\x46\x6a\x44\x2f\xfb\xd6\x35\x6a\x94\xe0\x4e\x67\xa3\x07\x77\x3c\xad\xd7\x48\x8d\xe8\x49\xa0\x61\xf0\x5a\x82\x3b\x9d\x46\xe3\xe1\xc4\xe0\x35\x18\xc9\x1c\x21\x1b\xa6\xc8\x56\xb8\x2e\x5e\xd3\x49\x70\x47\xa3\xa9\x3b\x32\x78\x4d\x8d\x64\x5f\x0e\x64\xc3\x60\xa3\x04\xd7\xca\x46\x7b\xb9\x06\x39\xd1\x10\x47\xfd\x82\x46\x57\x77\xa3\x51\xff\x76\x50\xfd\x0e\xa5\x71\xaf\xf0\x29\x23\xdd\xea\xbb\xd5\x09\x82\xd2\x40\x90\x48\xd7\x4e\x89\x42\xc5\xaa\xb9\x08\x40\xbe\xcc\xe7\xe7\x52\x8d\x39\x5e\x49\xc7\x5d\x5a\x5e\x2c\x28\x0a\xf0\x42\xa4\x43\xf2\xf8\x20\xd4\x80\xac\x31\xae\xe3\x40\x96\x74\xb0\xbf\xd9\x93\x25\x65\x22\xff\xef\x0a\x06\xf1\x3e\x15\x1c\xa9\x22\xbd\xd3\x3c\xb6\x6d\xbf\x79\x93\x7f\x0d\x80\x9f\xc3\xcb\x8f\x39\x8e\x92\x8a\xb5\xc1\x46\x58\x74\x2a\xcc\x9c\xfa\x51\x2d\xf1\xc1\x0b\x62\x9c\x7f\x8e\x38\x3c\xf3\xfa\x49\x9d\xf1\xad\xbc\xba\xb2\x74\x57\xb2\xff\xd3\x00\x24\x5d\x5a\xbf\x29\xf9\x6a\x27\x7d\xf0\x9e\x64\xb5\x26\x86\x80\xaa\xd9\xfb\xf7\xce\x96\x95\x83\x69\xbd\xb7\xf2\xe6\xcb\x7c\x8f\x4e\x25\x8f\x64\x27\xbc\x57\x73\xb4\x90\x23\xf6\x6c\xa6\x0b\x3f\x27\xaf\xcf\xb2\x91\x72\x7d\xb6\x8b\x80\x9a\xc5\x3c\x7a\xcf\xe8\x5a\x16\x4b\xcc\x8c\x90\x05\xa9\x38\xfc\xf9\x4a\x41\xd3\xc5\x2c\x48\x20\xc9\xe1\x89\xe5\xfe\x52\x69\xaf\xd5\xb8\xac\x66\xe6\xcb\x26\x65\x34\x59\x20\xfc\x17\xb2\x04\x8b\x25\x4f\x7e\xfe\x93\xd3\x10\x55\x9d\xe6\xbf\x03\xa8\x99\xca\x3a\x02\x50\x49\xf4\x3e\x7c\x46\xfd\xb6\xf5\xdf\xfa\x69\xe7\xc0\x44\x7c\x88\xc0\xa4\xb3\xcd\xe9\x89\x64\xe7\x1f\x27\x94\x3e\xce\x1a\x51\xea\xd2\xac\x83\x83\x07\xb3\x96\x4f\x18\x9e\x0b\xca\x36\x9d\x0e\x47\x99\x67\x5a\x66\x3d\x99\xc0\x19\xe6\xe0\xd0\x9e\x4f\xd6\x7a\x99\x24\x8c\x62\x71\x2e\x6d\x19\x47\xa8\xf0\x7c\xce\x57\x28\x7d\x21\xe9\xb0\xca\xb9\xab\x08\xf5\x7c\x9f\x86\x96\x33\x4e\xe0\x3a\x0d\x29\x84\x86\x2f\x28\x5b\xa9\xcc\xc5\x68\x90\xc9\x55\x08\x0b\x89\x4d\x84\x93\x73\x81\x85\x54\xde\x5b\xd2\x40\xda\x4e\xa5\x99\x01\x5e\x78\x71\x20\xd2\x20\x93\xcb\xd1\xd1\x95\xf4\xbf\xd3\x6d\x79\xf4\x26\x14\x22\xcc\x42\x6b\xc8\x46\xea\xe5\xa4\x33\x0b\x08\xb9\x82\x0b\xb9\xf0\xcc\x44\xff\x06\x00\x00\xff\xff\xd9\x02\x4a\xb6\xf4\x27\x00\x00")

func defines_html_bytes() ([]byte, error) {
	return bindata_read(
		_defines_html,
		"defines.html",
	)
}

func defines_html() (*asset, error) {
	bytes, err := defines_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "defines.html", size: 10228, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\x7b\x6f\xdb\xc6\xb2\xff\x3f\x9f\x62\x2f\x93\x16\xb7\x40\x45\x21\xb9\xbd\x38\x45\x6a\x0b\x50\x2c\x39\x11\xea\x87\x60\xc9\x3e\xa7\x7f\x15\x94\xb8\x92\x36\xa6\x48\x96\xbb\x92\xe3\x06\xfa\xee\x67\x66\x1f\x7c\x89\x2b\x51\xb4\xdc\xe3\x53\x24\x40\x4c\x72\x77\x76\x76\x76\x5e\xfb\x9b\x25\xed\xaf\x5f\x7d\x3a\x63\x21\x25\x8e\xba\xf2\xf7\xef\xd5\xcd\xef\xd3\x78\x95\x44\x0f\xdc\xd9\x6c\xbe\x7e\x7d\x60\x62\x41\xdc\xeb\xbb\xfe\xcd\xcd\xa0\xd7\xc7\x16\x17\x7f\xd0\x80\x53\xbc\xbe\x39\x23\xef\x4f\x49\x94\x10\xf7\xec\xa2\x3b\x1a\x5d\x75\x2f\xfb\xc4\x99\x06\x1e\x57\xa3\x13\x2f\x9c\x53\xf2\x86\xfd\x48\xde\x4c\xa3\x84\x22\xad\xdb\xf3\x84\xe7\x9e\x0d\x6f\xdd\x0b\xc6\xc5\x66\x73\x22\x12\x72\x4f\x1f\x4f\x81\x17\x92\xb8\x57\x9b\x4d\xe7\x44\xf8\x44\x32\x39\x75\x04\xfd\x22\x5a\x09\x9b\x2f\x04\x09\xa3\x87\xc4\x8b\x9d\x4e\x8e\xf2\xa4\x2d\xfc\x6a\x6a\xa7\x73\xc2\x63\x2f\x24\x28\xe2\x66\x93\x72\xbf\xe5\x34\x39\x43\x5a\x98\x25\xdf\x86\xac\x90\xbe\x73\x30\xc7\xd1\x23\xd7\x0c\x49\x27\xd7\xd6\x9c\xe1\xc0\x0f\x68\x59\x44\x6c\x2b\x71\x6c\x8b\x04\xba\x69\xe8\x4b\x73\xe4\x2e\xaf\x76\x1a\x56\x78\x93\x80\x3a\xa8\x76\xbc\x49\x85\xc2\x87\xb7\x84\x8b\x84\xc5\xb4\xa5\x68\x40\xe8\x05\xf5\x50\xf6\x04\x6f\x71\xc6\x05\xde\xec\xb0\x0c\x6a\x52\xad\x49\xd3\xac\x42\x06\xcb\xfc\x2e\x13\x7c\x1f\x07\xd0\xdc\xd3\x18\xa0\xa6\xf6\x72\x40\xdd\xe1\x9d\x5a\xde\x24\xf2\x1f\x41\x95\x82\x2e\xe3\xc0\x13\x3b\xe2\x81\xb8\xd2\xe3\x24\x3d\x5c\x51\x4b\x9d\xfd\x4a\xf7\x67\x93\x47\x41\xf9\xef\xc7\x0d\x29\x9f\xf1\xfb\x2c\x9c\x7a\xe7\x72\x8e\xed\x90\x42\x32\xb7\xc7\x92\x2b\x6f\x49\x7f\xa5\x8f\xc5\xd8\xca\x05\x94\xa6\xe3\xf7\x48\xf8\x69\x7c\x79\xb1\x1d\x5b\xdb\xd4\x89\x9d\x38\xef\xe5\x66\x40\x77\xed\xb1\xa0\x1e\x29\x38\x12\x28\xf5\xfb\x70\xc2\xe3\x5f\x4e\xf8\x2a\xde\x8a\x14\x43\x35\xa4\xc9\x94\x86\x22\x17\x2f\xa5\x9e\xcd\x26\x33\xbd\x64\x54\x63\xf6\x71\x24\xbc\x54\xd0\x26\x71\x66\x4c\x6e\x82\x6d\x8f\x51\x0f\x8b\x45\xed\x40\x6f\x2e\x32\xf3\x5f\xb0\xf0\x1e\x33\x4e\x16\x19\x48\x4d\x13\xa7\x43\xb2\x7f\x27\x1e\x59\x24\x74\x86\xda\xbb\x00\x87\x39\x1f\xb9\x9f\xe0\x11\xd2\x16\xfc\xcb\x14\xab\xbb\x4c\x4a\x23\xa4\xd3\xa3\x6b\x36\xd5\x31\x55\x41\xe8\x25\x54\xe4\xa8\x53\x65\x7b\x5b\xd1\xba\x5f\xa6\xcb\xa1\x55\x26\xe8\xca\xcd\x72\x19\xad\x42\x41\xfd\xb2\x57\x18\xc2\x43\x65\x22\x85\xa4\x5c\x94\xa9\x7b\xd7\x1d\x5c\x18\xb1\x8a\x53\xa9\x2e\xe3\x7b\xd2\xbd\xb5\x9a\xaa\x09\x73\x62\x3d\x4d\xa6\xdb\x51\xbf\x97\x6a\xaa\x38\x95\xec\x4a\x37\x24\x8c\x23\x23\x52\x25\x61\x5e\x53\x4f\x93\x69\x7c\x3d\xee\x5a\xf4\xa4\xba\xcc\xca\x65\x74\x55\xeb\x49\x13\xee\xd2\x93\x0e\xbe\x43\x53\x78\x21\x09\x37\xcd\xe3\x2c\x8c\xfc\x67\x4f\xe4\x6a\x92\x97\x9d\xc9\x07\xb3\x84\xd2\x9a\xa4\xab\x5a\xa9\x1c\xc9\x2c\xb9\x3c\xd7\xd5\x24\x99\x0f\xa4\x3e\x9f\x96\xcd\xb5\xe1\x1b\x63\xa7\x72\x06\xd4\x29\xd5\x92\x1f\x4d\x72\xdb\x1f\x82\x2a\xe5\xec\xa7\x93\x69\x60\x3f\x99\x8a\xcc\x26\x08\xa9\x18\x1b\xcd\xe2\x8b\x1d\x88\x93\x4c\x6d\x31\xcb\xa2\x67\x60\x83\x41\x6c\xe6\x16\xe2\xa6\x93\xb5\xd5\xf2\x7a\xa0\xed\xd1\x40\x78\x83\xb0\x36\xe9\xf5\x4a\xd4\xa1\xad\xc7\x31\xc7\xac\x89\x07\xb3\x12\x1e\x39\x1c\xfc\x0f\xc0\x23\x93\x99\xb7\xed\xb4\xdb\x10\x9c\x08\x26\x02\x7a\xea\x7c\x18\x8c\x47\x24\x06\x17\xe3\x74\x1a\x85\x3e\x40\xf3\xb0\x02\x98\x9f\xb0\xce\xe4\xa4\xcd\x3a\x31\xaf\x8d\xf1\xad\x13\x80\x96\x8e\x3b\x83\xc0\x80\x20\x1f\x7e\x1b\xf7\x47\x64\x19\xf9\xab\x20\x22\x3f\x7d\xb4\x2f\xe4\x03\x4e\xf3\xdd\x4f\x1f\x8f\x35\x8f\x75\x3d\x95\x13\x1d\x18\xb2\xec\xe9\x3b\x22\x9b\xd1\x24\x89\x92\x27\x86\xac\xe2\xf1\x2d\x66\x6d\xca\xfd\x4b\x83\xb6\x18\xaf\x66\xd7\xde\xf2\xc1\x06\xb1\x54\x0a\xd4\x63\xb2\x2e\x84\xa6\x95\x71\x93\xc0\x2c\x06\x63\x6d\xd6\x07\x87\x62\x3e\x8e\x9a\xc6\x62\xec\x4d\xef\xa9\x78\x62\x30\x6a\x26\xdf\xa2\xd1\xaa\xde\x6f\xe1\x58\x6f\x37\x7b\xbe\xa0\xdc\xe2\xff\xac\xa1\x59\x08\xab\x66\xb1\xb9\xa4\xcb\x43\x83\xf2\x47\xf2\x06\x46\x65\x91\x79\xd9\xbf\xdc\x8e\x49\xa0\x70\x7f\x65\x38\xb9\x89\xc8\xac\x65\x4f\x44\x20\xe1\x79\x9d\x22\x0e\x09\xb7\x4e\xe3\xb2\x76\x5d\x98\x99\x14\x50\xa3\x28\xc3\x81\x4f\x3e\x60\x03\x26\xcf\x71\x90\xed\x74\x50\x27\xbb\xfa\xab\x2a\xa9\xe3\x94\x50\xc6\x49\x9a\xb9\x58\xcc\x0f\xf4\xb0\x38\x89\xa6\x99\x7b\x0d\x47\x72\xa2\x6d\x17\x43\x32\x77\x38\xe8\xd9\x5e\xca\x38\x1d\x52\xa0\xda\x65\xfc\x94\xf2\x36\x4f\x99\x3f\x0e\xcc\x28\x38\x4d\xf6\x6f\x2b\xd9\xd4\x09\x8b\x12\x26\x1e\x6b\x92\x5f\x41\xf1\x5d\x93\x74\xc4\xfe\xac\x4b\x7a\x43\x39\xf3\xe5\x21\x45\x25\x39\x86\x09\x56\xf7\x86\x7c\xcc\x96\x15\x9c\xcd\xd1\x4c\x41\x1b\xe5\x4d\xb6\x49\xc4\xc4\xbc\xf9\x61\xf4\xbb\xa3\x1c\x46\x17\x8e\x1a\x8a\x27\xbf\xc3\x11\x38\x4f\xe5\x81\xa6\xea\x49\x4f\x28\xe1\x21\x3d\xce\xac\xa2\x7b\xc2\x71\x66\x59\xa2\x5b\xab\x44\xb7\x79\x89\x6e\xed\x12\xdd\x1e\x2e\x51\xe1\x70\x7c\x4b\xa2\x51\xff\xa6\xea\x80\x55\xf7\x98\x73\x33\x7c\xa8\x3c\xf3\x35\x74\xcd\x8f\xa1\xb7\xac\x76\x33\xb0\x59\x0d\x7a\x32\xab\x65\xf2\x54\xd2\x1d\xd1\x6a\x57\x83\xb3\x7e\xb5\x8e\x64\x8f\x59\xf6\xd5\xc0\x2a\x91\xa2\x3b\x9e\x8e\xee\x06\x37\xe3\x6a\x89\x64\x8f\x99\x04\x1f\x2c\x56\x53\x74\xc7\x93\xe8\xa6\x3f\xb2\x58\x0d\x7b\x52\x43\xc0\x83\xcd\xb3\x25\x5d\x13\xab\x99\x24\x58\x92\x68\x3c\xb8\xb4\x58\x4d\xf6\xa4\x6f\x0e\xe0\xc1\xa2\x23\x45\x77\x90\x8e\x76\xc6\x1a\xe6\x44\x8b\x1f\x75\x73\x12\x9d\x5d\x5f\x5e\x76\xaf\x7a\xdb\xef\xa1\x0c\xdd\x33\xbc\xca\xd0\x5b\x7d\x33\x9c\x20\xa2\x28\x10\x2c\x36\xd8\x49\x09\x94\x47\xcd\x79\x02\xe2\x43\x3e\x6f\x81\x18\x53\xba\x04\xc3\x9d\x3a\xde\x4a\x44\x24\xa0\x33\xa1\xbb\xa0\x7c\x10\xb2\x03\x30\xc6\xf9\x2a\x00\x48\x87\xfb\x9b\x3b\x5a\x44\x89\xc8\x3e\x57\xb0\xb0\xf7\x23\xc1\x89\x04\xe9\x8d\xe6\x71\x5d\xf7\xd5\xab\xec\x6d\x00\x5e\xf6\x2f\x7f\xc5\x69\xac\x10\x6b\x8d\x8d\x30\x6f\x54\x9c\x59\xdb\x51\x2e\xf1\xce\x0b\x56\x34\x7b\x1d\xb1\x7f\xe6\xf5\x5c\xd6\xf8\x4e\x86\xae\x1c\xd3\xa4\xf6\xff\x28\x40\x4e\xa7\xce\xff\x4b\xfe\x72\x27\xbd\xf3\xe6\x80\xd6\x44\x1f\xa9\x2a\xf6\xfe\x9d\xb3\xa5\x70\x50\xe3\xbd\xa5\x37\x5d\x64\x7b\xb4\xe6\x7c\x09\x8d\x38\xae\xa2\xb4\x80\x1e\xf7\xf6\xd6\x00\xbf\x4e\x86\xcf\xd2\x9e\x22\x3e\x2b\x53\x20\x66\xb1\xf7\x0e\x93\x68\x0d\x60\x29\xb1\x53\x00\x20\x15\xfb\x5f\x5f\x49\x52\xbd\x98\x19\x0b\x20\x38\x3c\xb1\xd8\x0d\x95\x76\x6a\x8d\x03\x9a\x99\x2e\xea\xc0\x68\x36\x23\xf4\x0f\xe2\x88\x64\x05\x71\xf2\xbf\x9f\x79\x14\x92\x6d\xa3\xf9\x3f\x20\xa9\x3d\x94\x8d\x07\x90\x02\xeb\x5d\xf4\x69\xe8\x1f\x8a\xff\xd6\xf3\x52\xc1\xc4\x7c\xf4\x40\xd5\x78\x48\xf5\xc4\xd2\xfa\xa7\x13\x82\x8d\xd3\x87\x58\x9b\x34\x6d\xe0\x68\xc1\xf4\xc9\x67\x09\x9d\x8a\x28\x79\x6c\x54\x1c\xa5\x96\x39\x30\xeb\x41\x02\x4f\x28\x47\x83\x9e\xf8\x6c\x6d\x96\xc9\xc2\x78\x25\x5a\xa0\xcb\x55\x4c\x72\xf7\x2d\xbe\x24\x7a\x80\x6a\x70\x8a\xb9\x2b\x4f\xea\xf9\x7e\x14\x3a\x9d\x1b\x45\x6e\xd2\x90\xa4\x30\xe4\xb3\x28\x59\xca\xcc\x95\x44\x41\xca\x57\x52\x38\x44\x3c\xc6\x54\xd5\x05\x0e\x91\x79\x6f\x11\x05\xa0\x3b\x99\x66\x7a\x74\xe6\xad\x02\xa1\x9d\x0c\x96\x63\xbc\x4b\xb5\xff\x60\x9e\xa1\xf4\x66\x11\x7a\x98\x43\xd6\x98\x8d\xe4\x60\xd5\x98\x3a\x04\xac\xa0\x0d\x0b\x4f\x55\x74\xf2\x3f\xbd\xeb\xb3\xf1\x6f\xc3\x3e\x59\x88\x65\xd0\x39\xd1\x3f\xa5\x15\x96\x54\x78\x64\xba\xf0\x12\x4e\x21\xd3\xae\xc4\xac\xf5\xb3\xa3\x5b\x17\x42\xc4\x2d\xfa\xc7\x8a\xad\x4f\x9d\x7f\xb5\x6e\xbb\xad\xb3\x68\x09\x91\xc6\x94\xef\x98\xe4\x3c\xe8\x9f\x52\x7f\x4e\xcd\x20\xf4\x0e\x70\x31\x46\x1f\x62\xd8\x0e\x72\x74\x0f\xcc\x17\x8b\x53\x5f\xbe\x03\x6d\xc9\x87\x1f\xc1\x0c\x4c\x30\x2f\x68\xf1\xa9\x17\xd0\xd3\xb7\xe8\x75\x78\xd8\x93\xa6\xc2\x4f\x11\x17\xa1\x4c\x29\x04\xee\x80\x0d\x18\x5f\x12\x9c\x04\x50\x6e\x80\x7a\x03\xb0\x0f\xcc\x60\x54\xcb\x96\xde\x9c\xb6\xe3\x70\xee\xa8\x6d\xdd\x69\xcf\xbc\x35\x12\xb8\xd8\x96\x1f\xc5\xc5\x63\x00\x96\xa1\xb4\x60\x96\xf6\x14\xf6\x00\x33\x14\xee\xdb\x2c\xf4\xe9\x17\x17\x5b\xc1\x29\xa6\x10\x1f\x22\x4f\xfd\xd9\x5b\x7b\xaa\xd5\x29\xab\x90\xf0\x64\x0a\x3c\x3e\xf3\x76\x82\x1a\x4c\x28\xdc\xbd\x73\xdf\xba\x6f\xff\x61\x1a\xdc\x25\x0b\xdd\xcf\x5c\xef\x75\x4b\x8f\x85\x8a\x5e\xda\xdf\x1d\x77\x3f\x7e\xec\xf7\x26\x2c\xdc\x6c\x80\x4e\xcb\xa1\x46\x98\x94\x01\x33\xb4\x97\x2c\xb8\xd7\x9d\xb2\xc3\x58\x5f\x49\x05\x37\xca\xc8\x2a\x70\x30\x16\x30\xfe\xa7\x01\xc3\xad\x50\x3b\xc9\x49\xe8\xa5\x21\x02\xb7\x13\x2f\x21\xea\xd2\xf2\x95\xeb\x99\xc7\x19\xfb\x42\xfd\x96\x88\x62\x87\x80\x6f\x53\x49\xcd\xe6\xe0\x0f\x18\x12\xf9\x40\x43\x93\xc3\x72\x28\x0c\x09\x56\xcc\x2f\x76\x6a\x66\x06\x8c\x55\x74\x4d\x20\x91\xfa\x12\x2c\xa2\xb0\x0b\xed\x02\xa9\x59\xb6\x71\xc3\x24\x12\x22\x5a\xea\x76\x11\xcd\xe7\xf2\xa4\x33\x8a\xa3\x35\xcc\xa0\x5b\x01\x14\xcf\x31\xd2\x16\xd8\x48\x66\xd1\x74\xc5\x73\x28\x43\x4a\x7b\xea\xbc\xf6\x19\x26\xc1\x49\xab\x38\x38\x75\xe1\x64\x15\x86\x2c\x9c\x13\xe3\x9e\xbd\xc1\x68\x7c\x33\xf8\xa0\xc1\x41\xc9\x63\x15\xe0\x93\x1a\x96\xe9\x04\x17\x53\xe6\x9f\x81\x43\x49\x67\x0c\xa4\x15\x01\xe0\x20\xf0\x62\x9e\x25\x6a\xd3\x40\xca\x04\xd9\xd0\x55\x2c\xd8\x92\xb6\x62\x40\xa0\xa1\xd8\x1a\xe8\x74\x54\x7f\xfe\xdc\x34\x1b\x95\xad\xe2\x56\x3e\xe7\x3e\x3d\x95\xe2\xad\x82\x9c\xa5\x8c\x0c\x70\x91\x81\x25\xf9\xb0\xd8\x29\x19\x53\x26\xbb\x94\xef\x60\x88\x3c\x03\x96\x0e\x08\xbc\xdd\x03\x2e\xba\x66\x40\x7b\x15\xd8\x25\x30\xb7\xe6\x1b\x25\xe0\x5e\xc5\x35\x37\x3e\x00\x4b\x40\x5a\x0e\x70\xe3\x28\xad\x24\x2d\x53\x9c\xd7\xe8\xec\x1d\xfc\xa9\x8c\xa9\x04\xcf\xf5\xc7\x90\x17\x5e\xe3\x59\x79\x65\xef\x1a\x52\xce\xeb\xb5\x02\x06\x19\x89\x5c\x88\xd1\x81\x96\xc6\x9b\x0a\xb6\xa6\x59\x81\xe4\x60\xea\xe5\xef\xdb\xed\x87\x87\x07\x17\x3c\x2a\x81\xff\xee\x34\x5a\xb6\x55\x16\x84\x14\x12\x50\x8f\x53\xde\xc6\x7d\x93\x8b\xf4\x8c\x1c\x7c\x8a\x43\x38\xa6\x0e\x0a\x30\x66\x34\xb8\xbe\x42\x07\xbd\x1e\x8d\xfb\x57\xe3\xb2\x18\xd2\xae\x99\xf3\xed\x0c\x60\xb4\x17\xa4\xe9\x44\x94\xbd\x0b\xd0\x09\x29\xb9\xa2\x59\x17\x92\x13\x35\xc8\x47\x44\x9a\x98\x07\xc6\x97\x8c\x73\x0d\x35\x26\x2b\x08\xe0\x74\xc3\x9d\x06\x11\xfa\xbc\xca\xb3\xaa\x4b\x07\xa2\x1e\xa5\xf9\x3a\xc4\x4b\x98\xd7\x5a\x30\xdf\xa7\x90\x3b\x25\x26\xeb\x7c\x8f\x8e\xcb\x7f\x39\x69\xab\x81\xb9\xd8\x53\x13\x43\x27\xf7\x70\xb7\xea\xe2\x63\xd1\xbf\x0b\x3f\xc1\x1b\xb2\x65\xcb\xa4\x57\xad\x1c\x12\x03\x28\x80\xa4\x80\x89\x91\xd3\x35\xe8\xe5\x11\x90\x47\x41\x0b\xa0\x9f\x62\xc3\x24\x88\xa6\xf7\xa8\xb2\xd6\xd2\x6f\xbd\x33\x37\xd1\x6c\x06\x1b\x48\xeb\x6d\x91\x16\xe4\xa3\x01\xb6\x05\xde\x84\x06\xe5\x80\x6e\xa9\x56\x49\x24\x73\x2a\xa6\xa7\x89\x08\x5b\x72\x0a\x07\x12\x1d\xa4\x35\x7f\xb5\x5c\x3e\x2e\xe9\x12\x24\x9f\xb1\xf9\x74\x41\xa7\xf7\x93\xe8\x8b\xec\x6c\x71\xdc\x0c\x5f\xa7\x9d\xd2\x05\x8d\x94\xc0\x46\xa5\x69\xc5\x2d\x45\xbf\x33\x2f\xe0\x25\xf8\x7b\x26\xb7\x14\xf7\x13\x40\x41\xc5\xe8\xb2\x7f\x09\x18\x85\x28\xc7\x36\x1b\x93\x09\x0c\x98\xce\xe9\x5c\xd2\xa5\xc4\x85\x9e\x01\x50\x32\x41\x5a\x44\x2d\x2f\x3b\xeb\x51\x6e\x92\x3d\x6f\xfb\x04\x51\x4f\xe8\xf7\xa8\xac\x12\x5e\x7b\x1a\x4b\xe3\x20\x99\x02\x0d\xdf\x1d\xb5\x82\x45\x59\xa9\x24\x8a\x79\xb6\x9b\x23\xa4\x4c\x0f\x31\xa2\x84\xfd\x89\x1e\x18\xb4\x64\xf3\x04\xca\x0d\x34\x90\xdc\x02\x65\x53\xd1\x7d\x24\x1c\x55\xe0\x15\x1d\x95\x96\x76\x63\xb4\xb1\x42\xc4\xe9\x1d\xe0\xe1\xd2\x56\xaa\x42\x89\x97\x7d\x10\x46\xc8\x51\x06\x26\x14\x7d\xb2\xe4\x5e\x8e\x51\x7b\x49\xb9\x9d\x4f\x72\xb1\xa9\x6d\x76\xcd\x60\x90\xf1\x1e\x07\x1c\xfd\xb3\x3b\xb4\xfb\x1e\x28\x97\xf2\x07\x2c\x24\xed\x02\x51\x82\x14\x78\xce\xa0\xa5\xb2\xe4\xc9\x0a\xe5\x9a\x0a\x32\x2f\x94\x2e\x17\xc0\xca\x96\x3a\xcf\xd4\x2b\xb9\x0a\xcf\x4c\x89\x53\x94\xe1\x01\xaa\xb3\xec\xbe\x35\x83\x53\x87\x25\xcb\x79\x57\x9a\x20\x81\xad\x29\xfe\xf6\xbc\xb1\xd2\xe5\xa3\x5b\x06\x08\x16\x35\x15\xd2\xdd\xff\x3d\x4f\x7a\x63\x33\x7b\x76\x33\x7d\x47\x4b\x6e\x83\x73\xbb\x7f\xb1\x59\x86\x5e\xf4\xc0\xb1\x37\x19\x9c\xbb\x63\xdc\xa2\x35\x2c\x2c\x25\xbc\x6d\xe1\x5f\x7a\xbe\x4b\x55\xda\x24\xdd\x49\xf5\x35\xc9\x76\xff\xb1\xd4\x06\x32\x32\x59\x40\x1e\x27\x87\xed\xf6\xb3\xfe\x17\xf0\x77\x1f\x43\x4c\xf9\x19\xa0\x1e\x79\x70\x5a\xe5\x69\x16\x81\x4a\x0e\xa8\x38\x22\x06\x1e\x9c\x6f\x36\xc7\xce\x6b\xc8\xb3\x59\x5a\xdb\x82\xf3\x98\x7f\x78\x01\xbf\xb3\x59\xcb\x9c\x01\x29\xc3\x79\x13\xf4\xbf\xb4\xa0\x38\x1f\x77\x3f\x8c\x5c\x76\x3e\xec\x9e\xfd\xda\x1f\x8f\xdc\x5b\x86\xaf\x62\x73\x48\xfa\x75\xfa\x2d\x87\xd3\x19\xaa\x9b\x3c\x5c\x3f\x74\x9e\xfe\xcd\xcd\xf5\x4d\xf5\x34\xea\x6b\x2e\xa7\x23\xcf\x00\x77\x4f\xa2\x73\xc7\x9e\xb9\xe4\x97\xa1\x95\x53\xc9\x6f\x38\x9d\xce\x07\xbc\x54\x42\xfa\x2c\x4e\xcd\xe2\x73\x32\xc0\x7c\xa4\xbc\x6f\xe4\xfd\xb0\x52\xa9\x3f\x54\xb9\xaa\xca\x6c\xb2\xbb\xb0\xab\xe8\x75\xa9\xa2\xc4\x6a\x2b\x9d\x11\x0f\xb5\x6c\xba\x61\xa5\x8b\xab\xb3\x6d\x95\x3f\xe5\xda\xde\xbd\x32\x9d\x69\x4b\x36\x51\x59\xce\x3f\x8e\xa8\x31\xcd\xb5\x9e\xc2\xca\x2e\x9a\xd3\x97\x5a\x58\x3d\x75\x15\xbe\x43\xdd\xa5\x2d\xe5\x8c\x4d\x94\x95\x39\xf8\x11\x75\xa5\x98\xd6\x53\x55\x29\xc2\x72\x9a\x92\x8b\xaa\xa7\xa8\xfc\x47\xf6\xcd\x30\xd1\xbb\xe7\xc1\x44\xd3\x78\x65\x07\x45\x69\xe7\xd1\x50\xd1\xd9\xf0\x16\xec\x54\x8d\x8a\x60\x36\xa7\x03\x04\x55\xe8\xa7\x42\xcc\x97\x0e\x7f\x32\xe5\x35\xc1\x3f\x4a\x51\x7f\xdf\x72\x4f\x1a\xfb\x2f\x86\x4a\x4a\xa7\x56\xac\xb4\x4b\x24\x2b\x58\x02\x9e\xc7\x47\x4b\x92\xe9\x31\xaa\x40\x5c\x52\xc3\x2a\x50\x2b\xab\xb2\x0a\x04\xb6\x75\xd2\x5e\xfa\x37\x05\x5e\x54\x15\xe8\xef\xa8\x02\xfd\x63\x57\x81\xbd\x73\x6b\xba\xf3\x2b\x8b\xc0\xde\xce\x22\x70\x5b\xf6\x97\x9e\x05\xfd\xa7\x14\x81\xbd\x6f\x45\x60\xbd\xcc\xd6\xdb\x59\x04\xfa\x0d\x8a\xc0\xde\x33\x14\x81\xbd\xe7\x2c\x02\xfd\xdd\xc5\x59\x4f\xc1\x39\xff\x7c\x70\x75\xdd\xab\xac\x98\xcc\x2f\xaa\xe2\x2f\x25\xe0\xb5\xaa\x38\xf3\x6b\x15\x67\xe9\x5c\xb6\xe2\xcc\xaf\x5f\x9c\xa5\x52\xe5\x44\xd8\x87\x9d\xab\xd6\x6a\x03\xcf\xbd\xda\xe0\xb9\xcc\x75\x27\x7a\xb6\xaa\x3b\x7b\xcd\xa9\x17\x56\x67\x23\x29\xfd\x9a\xb5\xbd\xd0\xf0\x4b\x85\xc6\x21\xca\xda\x5f\x68\x34\xd0\x55\x8d\x42\xc3\xe6\x2d\x39\x4d\xd5\x2e\x34\xfc\x03\x0a\x8d\x83\xde\xeb\xd5\x7a\x4b\xf5\x73\xe9\x2d\xd5\x33\x95\x2c\x31\xb7\x6f\xe0\xa6\xef\x68\x1b\xf8\x70\x64\xdd\xc0\xf1\xc5\xee\x30\x89\xa6\x94\x73\x1d\xc4\xa5\xdd\x7a\x5b\xd0\x97\xbe\x5b\xa7\xea\x6b\xb2\x5b\x4b\x55\xd5\xda\xad\xd5\x3b\x75\x5b\xb5\x52\xfc\x53\x22\x8d\x77\x9b\xe1\xe8\xe0\xdd\xe6\xa5\x15\x49\xe8\x61\x0d\x90\x44\x00\x1e\x49\x6c\x70\xc2\x6a\xc8\xe1\x28\x8c\x44\x8f\x4e\x13\xea\x49\x18\xb1\x0b\x51\xc4\x1c\xe7\xb0\x0a\xd7\xaa\x96\x6b\x89\x7f\xe7\xad\x99\x5c\x19\xd4\xd9\x2d\x16\x4e\x51\x17\xec\x0c\x47\x71\xb0\xe2\x63\x70\x36\x2b\xd4\xa9\x2c\xaa\x62\xde\xb4\xa6\x92\x11\x52\x5d\x52\xc5\xb5\xf2\xbb\xf9\x55\x9d\xbf\x69\x66\x5f\xcf\xed\x99\xdd\xf4\x1d\x2d\xb3\xdf\x7d\xb4\x66\x76\xfc\x28\x47\x7f\xac\x4b\xe6\x41\x34\xc1\xaf\x0f\x85\x27\x56\x95\x59\x7e\x5b\xe8\x97\x9e\xe5\x53\x55\x36\xc9\xf2\x52\x6d\x2f\x27\xcb\xdf\x7d\xfc\xaf\xcf\xf2\xeb\x7d\xf5\xe2\x21\xc9\x69\x3d\x6f\x9a\x9c\xa4\x61\xab\x93\xd3\x7a\x5e\x27\x39\xa5\x9f\x8c\xd7\x4e\x4e\x3b\xbf\x96\xed\xbc\x5a\x7b\x09\x41\x29\xc9\xa9\xf9\x58\x6d\xb3\xf9\xe5\x55\xee\xc3\x55\xfd\xad\xb7\xfc\x54\xf9\xdf\x01\x00\x00\xff\xff\x6c\xf7\x8b\x70\xb4\x54\x00\x00")

func index_html_bytes() ([]byte, error) {
	return bindata_read(
		_index_html,
		"index.html",
	)
}

func index_html() (*asset, error) {
	bytes, err := index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "index.html", size: 21684, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
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
	"defines.html": defines_html,
	"index.html": index_html,
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

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"defines.html": &_bintree_t{defines_html, map[string]*_bintree_t{
	}},
	"index.html": &_bintree_t{index_html, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
