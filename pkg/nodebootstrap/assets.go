// Code generated by go-bindata.
// sources:
// assets/10-eksclt.al2.conf
// assets/bootstrap.al2.sh
// assets/bootstrap.ubuntu.sh
// assets/kubelet-config.json
// DO NOT EDIT!

package nodebootstrap

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

var __10EkscltAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x6b\xdb\x4e\x10\xc5\xef\xfa\x14\x0b\xc9\xe1\xff\x07\xaf\x14\x3b\x6e\x0e\x01\x1d\x4c\xad\x84\x80\x6b\x87\x28\xa1\x85\xb6\x98\xf5\xee\xd8\x4c\xbd\x9a\x5d\x66\x57\xb6\x53\x93\xef\x5e\x64\x4b\xa9\x4b\x42\xe9\x4d\x9a\xdf\xe8\xbd\x99\x37\x3a\x13\xb0\x0e\x3a\x5a\x19\x3c\x68\x5c\xa2\x16\xe1\x39\x44\xa8\x8c\x30\xec\xbc\x44\x12\x35\x61\x14\x4b\xc7\x62\x5d\x2f\xc0\x42\xec\x1d\x5e\x46\x95\xfa\xe9\x48\x4c\x90\xea\x9d\x18\x88\xff\x46\x93\xc1\xff\x49\xf2\xb5\x04\xde\xa0\x86\xef\xc9\x99\x98\x38\xad\xac\xa8\x20\x2a\xa3\xa2\x12\x5e\xb1\xaa\x20\x02\x87\x6b\xf1\x50\xdc\xde\xcd\xa6\x3d\x31\xfa\x5c\xce\xc7\xc5\xcd\xe8\x69\xf2\x38\x3f\xd6\x92\x82\x36\xc8\x8e\x2a\xa0\x78\x83\x16\xf2\x0c\xa2\xce\x8e\x23\x66\x9d\x56\x0a\xb4\x49\xce\xc4\xad\x75\x0b\x65\x85\x22\x23\x42\x54\x11\xf5\x1f\x1e\x9f\x46\x5f\xe6\xf7\xb3\x71\xd9\x13\x1f\x27\x4f\xe5\x63\xf1\x30\x1f\x4f\xcb\xbf\xca\xb7\xfb\xb5\xea\xc7\xf1\xc9\x91\x7c\x47\x7c\x3a\x1b\x17\xf3\xbb\xfb\x7f\x92\xb3\x8d\xd0\x41\x34\x29\x76\xa0\xcb\xa8\x38\xe6\x27\x8f\x59\x1d\x38\x5b\x20\x75\x1f\x88\x6f\x89\x10\x52\x92\x33\x20\xd1\xe7\xe7\xfb\xd6\xec\xa5\x05\xda\xd6\x21\x02\x4b\x43\x21\x3f\xdf\x9f\x2c\xd7\x35\x54\x6a\x27\xbd\x33\x0d\xed\x42\xe8\x90\xb2\xd6\x6d\xa5\x67\xdc\xa0\x85\x15\x98\x3c\x72\x0d\x2d\xf3\xce\x48\xa4\x25\x2b\xa9\x1d\x45\x85\x04\x2c\xb1\x52\x2b\xc8\xaf\x2e\x06\xc3\x8b\x7e\x7f\x78\x39\xfc\x30\x48\xcd\x9a\x53\xd0\x9c\x9e\xef\xdf\x5e\xef\x25\x55\x87\xdf\x42\x6d\x43\xaa\x5d\xd5\x24\x91\x79\x55\x07\x90\xaa\x32\x57\xc3\xeb\xcb\xb4\xff\xba\x84\xab\x8d\xf4\xec\x36\x68\x80\x73\xb5\x0d\x1d\x20\x94\x0b\x24\x69\x90\xf3\xcc\xf9\x98\x69\xc2\x26\x9d\x13\xac\x1d\x2d\x8f\xbc\x49\xbb\xe1\x04\x31\x35\x5d\xc7\xeb\xf0\x5c\x53\xc4\x0a\x72\xe3\xf4\x1a\xb8\x8b\x15\xe2\xd6\xf1\x5a\x7a\x5b\xaf\x90\x72\x4d\xd8\x02\x86\x15\x1e\x72\x6d\x82\x3f\xcd\xa5\x39\x4b\x63\x89\xab\x37\xe7\x3d\x96\xd3\x67\x55\xd9\xdf\xee\xef\x35\x5a\x88\x2d\x4a\x7f\x04\x47\xc9\xaf\x00\x00\x00\xff\xff\xa2\x51\x55\xb0\x76\x03\x00\x00")

func _10EkscltAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		__10EkscltAl2Conf,
		"10-eksclt.al2.conf",
	)
}

func _10EkscltAl2Conf() (*asset, error) {
	bytes, err := _10EkscltAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10-eksclt.al2.conf", size: 886, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcb\x31\x0e\xc2\x30\x0c\x46\xe1\x3d\xa7\x30\x85\x01\x86\x34\x27\x80\x09\x06\x16\xe0\x06\xc8\x49\x7f\x29\x55\x9d\x44\xaa\x5d\x24\x6e\xcf\x82\xaa\xae\xdf\xd3\xdb\xef\x42\x1c\x6b\xd0\x4c\x1e\x8b\x73\x48\xb9\x51\xf7\x78\x5e\x6f\xef\xfb\xeb\x7c\x38\xe6\xa6\x56\xb9\x80\xfc\x78\xea\xe8\x42\x01\x96\x02\x26\x4d\x26\x61\x5a\x22\x04\xd6\x4b\x4b\x2c\x3d\xea\xc7\x39\xfd\xaa\xa1\x24\x13\x1a\x18\xa5\x55\x3f\x43\x1a\x0f\x1b\x47\xe5\x28\xa0\xff\xbb\x09\x6a\x3c\xdb\xea\xbf\x00\x00\x00\xff\xff\xef\x6e\xff\x33\x97\x00\x00\x00")

func bootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAl2Sh,
		"bootstrap.al2.sh",
	)
}

func bootstrapAl2Sh() (*asset, error) {
	bytes, err := bootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.al2.sh", size: 151, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapUbuntuSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\x51\x6f\xdb\x36\x10\xc7\xdf\xf5\x29\x6e\x6e\x80\x25\x40\x29\x27\x69\xd6\x87\x02\x1a\xe6\x35\x5e\x51\x34\x8b\x0b\xc7\x41\x0b\x0c\x83\x71\x26\x4f\x16\x61\x8a\x27\x90\x27\x6b\x69\x90\xef\x3e\xd0\xb2\x5c\x27\xc0\x56\xf8\xc5\xe4\xff\x77\xa7\xe3\xdd\x9f\x7c\xf5\xd3\x78\x65\xfd\x78\x85\xb1\x02\x45\x6d\x96\x91\xae\x18\x46\xb7\xb3\xeb\xe9\xf2\xe3\xe7\xe2\xe4\xb4\xe2\x28\x1e\x6b\x02\x65\xcf\x46\xf0\x2b\x8c\x49\xf4\x98\x36\x51\x8b\x1b\x6f\xda\x15\x39\x92\xdc\xb1\x46\x97\x93\xdf\x66\x59\xf4\xd8\x00\x3a\x8b\x11\xf6\xaa\xa2\x4d\xcc\xf7\xff\x87\xbd\x97\x98\x16\x77\xc0\xb4\xb8\x61\xaf\xc7\xa2\x70\x73\x9c\x2c\x8b\x0f\x51\xa8\x4e\x5c\xa0\x48\xa2\x4a\xb4\x8e\x4c\x96\x9d\x66\x00\xaf\x60\x31\xbb\x9e\xbd\x03\xa9\x28\x12\xc4\x8a\x5b\x67\x60\x45\xe0\x98\x37\x64\x00\x05\x68\x4b\xe1\x01\xc4\xd6\x34\x24\x85\x28\x18\x24\x42\xdb\xbc\xde\x65\xe8\x2a\xab\x2b\xb0\x11\xba\x0a\x05\x3a\x02\xc3\x60\x3d\x4c\x6e\x2e\xe1\xf4\xa0\xad\x30\x92\x01\xf6\xd0\x38\xb4\x1e\xfa\x9a\x4c\x9f\x00\xbd\x81\x9a\xd0\x0b\x08\xa7\x8f\x37\x1c\x04\x57\x8e\xd2\xb2\xe6\x28\x03\x0d\xc6\x46\x09\x1c\xcf\x5e\xc3\xaa\x15\xb0\xf2\x73\xdc\xc5\x7b\x16\xd0\x8e\x30\x40\xc5\x5d\x0a\x72\x8c\x66\x7f\xa4\x32\x70\xfd\xbd\xf0\xd4\x9f\xce\x4a\xc5\xad\x40\x85\x5b\xeb\xd7\xbb\x04\xc2\xa0\xdb\x28\x5c\xdb\x48\x29\xae\x07\xad\x44\x72\x65\x06\x10\xb9\x0d\x9a\x7e\x30\xca\xff\xc5\xfe\x13\xa8\x49\xd0\xa0\x60\xef\x06\x80\xd2\xe1\x3a\x16\x69\x32\x00\x23\x34\x26\x50\x8c\xc5\x79\xbe\xfb\x8d\xfa\x5d\xcf\x86\x94\x6d\x8a\x93\xc7\xbd\xeb\x9e\xf6\x82\x76\x6d\x14\x0a\xca\xf8\x58\x9c\x3c\xbe\xbf\xb9\xbf\x5b\x4c\xe7\xcb\xeb\xdb\xbb\x01\xa8\xf1\x1f\xd5\xb0\x49\xea\x9f\x93\xaf\xcb\xcf\xb3\xeb\x83\x84\xad\x54\xe4\xc5\x6a\x14\xcb\x5e\x09\x6f\xc8\xab\x8e\x56\x15\xf3\xa6\x90\xd0\xd2\x11\xc7\xc1\x7e\xeb\xb1\x9a\x0d\x15\x5f\x7a\x6a\x00\x9c\xe3\x4e\x35\xc1\x6e\xad\xa3\x35\x99\xe3\xe0\x86\x8d\xb2\xbe\x0c\xa8\x34\x7b\x41\xeb\x29\x28\x5b\xe3\x9a\x8a\xb7\xe7\x97\x57\xe7\x17\x17\x57\x6f\xae\x7e\xb9\xcc\xcd\x26\xe4\xa4\x43\x7e\xf2\x38\xf9\x72\xb7\xbc\x9e\xfe\x31\xb9\xbf\x59\x2c\xe7\xd3\x0f\x1f\x67\xb7\x4f\x39\xd6\xf8\x8d\x3d\x76\x31\xd7\x5c\xa7\x36\x8e\x1b\x6c\x23\x29\xac\xcd\xdb\xab\x77\x6f\xf2\x8b\x43\x37\xb8\x35\xaa\x09\xbc\xb5\x86\x42\x81\x5d\x7c\xd9\x26\xae\xd1\xfa\x62\xbf\xec\x27\x39\x20\xde\xaa\x95\xf5\xca\xd8\x50\x8c\xb9\x91\xb1\xf6\x36\xdd\xfb\x23\x59\xb3\x2f\x7b\x3d\x4d\x33\xe9\x9e\x24\x37\x03\x71\x38\x5f\x68\x7d\xba\x3b\x85\x61\xbd\xa1\x30\x8c\x90\xa4\xe3\xb0\x51\x8d\x6b\xd7\xa9\x04\x6f\x87\xb8\x75\xe0\xb6\x51\x26\xd8\x2d\x85\xa2\x5f\x95\x43\xe1\x81\xd6\x76\x57\x79\x72\xc0\x71\x5f\x77\xd7\x9f\x7d\x69\xd7\xc5\x4b\xf3\xf5\xdb\xf9\x03\xd6\xc3\xd9\x4a\x42\x69\x03\xa9\x35\x0a\xc5\x62\xce\x82\x42\x9f\x7a\x9b\xde\x51\xd8\x52\x78\x4f\x41\x6c\x99\x9c\xf0\xec\x23\xe8\xd9\x3f\xd4\xdc\x46\x95\x3c\x50\x94\xe8\x22\x1d\x3a\x6a\xc9\x8b\xd2\xa8\x4a\xeb\xe8\x59\x0d\x1a\x73\x1d\x24\x71\x67\xc9\xde\xfd\x03\xf5\xfd\x61\x4b\xef\x13\x8c\x4e\x1e\x77\xb6\xff\xeb\xb7\xbf\x9f\x46\x69\xf5\xe9\xfe\xf7\xe9\xcd\x74\xb1\x9c\x7e\x5d\xcc\x27\xcb\xc9\xfc\x43\xb2\xe9\x59\x36\x3c\x6f\x18\x9e\xc5\x67\xff\x06\x00\x00\xff\xff\x6a\x77\xb9\x7b\x94\x05\x00\x00")

func bootstrapUbuntuShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapUbuntuSh,
		"bootstrap.ubuntu.sh",
	)
}

func bootstrapUbuntuSh() (*asset, error) {
	bytes, err := bootstrapUbuntuShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.ubuntu.sh", size: 1428, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubeletConfigJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x4d\x6f\xf2\x30\x10\x84\xef\xfc\x8a\xc8\x67\x94\x84\xf7\x15\x52\xcb\x8d\x82\xda\x43\x39\x15\xda\x9e\x1d\x67\x03\x56\x1c\x2f\x5a\xaf\xe9\x97\xf8\xef\x95\x3f\x5a\x0a\x3d\x54\x39\x79\xf4\xec\x78\x36\xe3\x8f\x51\x51\x88\x5e\xdb\x56\xcc\x0a\x71\xef\x1b\x30\xc0\x0b\xb4\x9d\xde\x7a\x92\xac\xd1\x8a\x71\x20\xe4\x5e\x3f\x01\xb9\x70\x9e\x15\xa2\x4f\x5c\xa9\x22\x58\xf6\x57\xae\xd4\x58\x1d\x26\x0d\xb0\x9c\xe4\x81\xb6\x25\x70\x2e\xd0\x75\x19\xbf\xac\x7b\xde\x81\x65\xad\x92\xf9\xac\x08\x01\x82\x6e\xd1\xbe\x0d\xe8\xdd\xb7\x54\x14\x02\xac\x6c\x0c\x84\x68\x9d\x34\x0e\xa2\x7c\x1c\xa7\x81\x17\x68\x76\x88\xfd\x4f\x5c\x49\xb5\x83\xcd\x66\x15\x2e\xfd\x37\xd4\x4e\x8c\x7f\x1b\x31\xf9\x73\x9f\xd7\x69\x7d\x7d\x66\x62\x34\x58\x5e\xcc\x6f\xb5\x81\x60\x54\x01\xab\x0a\x7a\xa7\xd8\x54\x4a\x96\x8a\x58\xa4\xf9\x51\xf6\x88\x2b\x21\xe9\xf7\x8b\x8d\x06\x6c\xa3\xc1\x73\x0e\xfa\x47\xee\x79\x76\x81\x36\x6f\x30\x3d\xdb\x20\x32\x8f\x56\x5e\x52\xff\x6b\x77\x99\x47\x19\xef\x18\x68\x89\x83\xd4\xb1\xae\x2c\x94\x06\x95\x34\xa9\x06\xb5\x25\xf4\xfb\x25\xe9\x03\x50\x44\xe2\xb9\x4b\x17\x8a\x0e\x24\x7b\x82\x3b\xc9\x70\xea\x43\x3c\x20\x4b\x86\xfc\x46\xd6\x40\x07\xa0\x05\x10\xeb\x2e\x94\x09\xa7\x7f\x9b\x52\xb8\x08\x6c\x56\xeb\x1b\x44\x76\x4c\x72\xff\x45\x1c\x3f\x03\x00\x00\xff\xff\xbb\x62\xda\x19\x74\x02\x00\x00")

func kubeletConfigJsonBytes() ([]byte, error) {
	return bindataRead(
		_kubeletConfigJson,
		"kubelet-config.json",
	)
}

func kubeletConfigJson() (*asset, error) {
	bytes, err := kubeletConfigJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubelet-config.json", size: 628, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"10-eksclt.al2.conf": _10EkscltAl2Conf,
	"bootstrap.al2.sh": bootstrapAl2Sh,
	"bootstrap.ubuntu.sh": bootstrapUbuntuSh,
	"kubelet-config.json": kubeletConfigJson,
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
	"10-eksclt.al2.conf": &bintree{_10EkscltAl2Conf, map[string]*bintree{}},
	"bootstrap.al2.sh": &bintree{bootstrapAl2Sh, map[string]*bintree{}},
	"bootstrap.ubuntu.sh": &bintree{bootstrapUbuntuSh, map[string]*bintree{}},
	"kubelet-config.json": &bintree{kubeletConfigJson, map[string]*bintree{}},
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

