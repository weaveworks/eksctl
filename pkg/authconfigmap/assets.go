// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/emr-containers-rbac.yaml (1.468kB)
// assets/ssm-containers-rbac.yaml (953B)

package authconfigmap

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _emrContainersRbacYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\x3f\x8f\xdb\x30\x0c\xc5\x77\x7d\x0a\x21\xf3\x39\x87\x6e\x85\xc7\x2e\xdd\x0f\x68\x97\xa2\x03\x25\xbf\xf8\x74\x27\x8b\x02\x29\xa7\x7f\x3e\x7d\x21\x39\xbe\x2b\x5a\xa7\x48\x8a\xa0\x9d\x44\x93\xf6\xd3\x8f\x0f\x26\x29\x87\x8f\x10\x0d\x9c\x7a\x2b\x8e\xfc\x9e\xe6\xf2\xc8\x12\xbe\x53\x09\x9c\xf6\xcf\x6f\x75\x1f\xf8\xfe\xf8\xc6\x3c\x87\x34\xf4\xf6\x81\x23\xcc\x84\x42\x03\x15\xea\x8d\xb5\x89\x26\xf4\x16\x93\x74\x9e\x53\xa1\x90\x20\x6a\x64\x8e\xd0\x5a\xed\x2c\xe5\xf0\x5e\x78\xce\xda\xdb\x4f\xbb\xdd\x67\x63\xad\xb5\x02\xe5\x59\x3c\x5a\xae\x0a\x68\x26\x0f\x3d\x55\x8f\x10\xd7\x2a\x23\x4a\x4b\x5d\x20\xa2\x90\x63\xf0\x20\xef\x79\x4e\x45\x77\x77\x76\x4d\xb5\xd8\x73\x3a\x84\x71\xa2\xdc\x9e\x70\xc4\xe9\x9d\xcc\xc3\xcb\x79\x1f\x79\xdc\x22\xb8\xb3\xbb\x18\xb4\x9d\x5f\xa8\xf8\xc7\x1a\x0c\x50\x2f\xc1\xa1\x69\x0b\xa8\xb4\x08\x43\x28\x4b\x35\x62\xc9\x2c\x91\xe7\x18\xe1\xab\x9d\x35\x47\x29\x71\x39\x7d\x91\x57\xc1\x48\x0e\xf1\xf2\x66\xbd\xa0\xfc\x66\xd7\x2b\x48\x7e\xe5\x5c\x49\x16\xf4\x8d\x0b\x28\x67\xdd\xbc\xa4\x32\x1e\xe6\xa8\x58\xac\x1a\x90\x23\x7f\x9b\x9a\x73\x37\x76\xe9\x3a\x47\xdc\x4b\x27\xbf\x10\x3f\xb1\xfb\xcf\x68\xf8\x5a\x90\xea\x28\x6d\x3a\x1a\xd2\x28\x50\xdd\xfe\xcf\xff\x1d\xe4\xd9\x21\xdf\x62\x16\x8e\xcb\x04\xd5\xc0\x85\x34\x84\x34\xde\x9c\xff\xea\x29\x31\x5d\xd7\x99\xbf\x58\x5c\xef\x96\x06\x2e\xd8\x5f\x1c\xf1\x80\x43\xad\xaf\xde\xfd\xe1\x12\x63\xed\x4f\xcb\xf1\x8c\xa4\xce\xee\x09\xbe\x68\x6f\xba\xab\x34\x3f\x28\xe4\x9c\xe6\x8f\x00\x00\x00\xff\xff\x8b\x95\x6b\x36\xbc\x05\x00\x00")

func emrContainersRbacYamlBytes() ([]byte, error) {
	return bindataRead(
		_emrContainersRbacYaml,
		"emr-containers-rbac.yaml",
	)
}

func emrContainersRbacYaml() (*asset, error) {
	bytes, err := emrContainersRbacYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "emr-containers-rbac.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xba, 0xbc, 0x51, 0xe8, 0x47, 0x23, 0x61, 0x8, 0xb1, 0x55, 0x1d, 0xce, 0x13, 0x1e, 0xda, 0x59, 0x4e, 0x6d, 0xda, 0x2b, 0xc5, 0x4, 0x3d, 0x7a, 0x8, 0xce, 0x29, 0xdf, 0x28, 0xab, 0x7d, 0x60}}
	return a, nil
}

var _ssmContainersRbacYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x53\xb1\x6e\xe3\x30\x0c\xdd\xf5\x15\x82\xe7\x73\x82\xdb\x0e\x1e\xaf\x43\xf7\x00\xed\x52\x64\x60\x24\xc6\x61\x23\x8b\x02\x29\xb9\x68\xbf\xbe\x90\x92\x14\x05\x92\x14\x2d\x0a\xb4\x13\x9f\x4d\xf2\xbd\x47\x9a\xee\xfb\xde\x40\xa2\x7b\x14\x25\x8e\x83\x95\x0d\xb8\x05\x94\xbc\x63\xa1\x17\xc8\xc4\x71\xb1\xff\xa7\x0b\xe2\xe5\xfc\xd7\xec\x29\xfa\xc1\xde\x84\xa2\x19\x65\xc5\x01\xcd\x84\x19\x3c\x64\x18\x8c\xb5\x11\x26\x1c\xac\xea\xd4\x3b\x8e\x19\x28\xa2\xa8\x91\x12\x50\x6b\xb6\xb7\x90\xe8\x56\xb8\x24\x1d\xec\x43\xd7\xad\x8d\xb5\xd6\x0a\x2a\x17\x71\xd8\xde\x29\xca\x4c\x0e\xb5\xfb\x63\x3b\xc7\x71\x4b\xe3\x04\xa9\x3d\xe1\x8c\x31\x37\x94\xd8\xbf\xc5\x65\xe0\xb1\xe2\xc8\xfe\xd0\xa4\xe8\x04\x0f\x75\xd5\x8c\x26\xa8\x6c\x07\xa5\x19\x65\xd3\x54\x46\xcc\xb5\x20\x90\xb6\xf8\x04\xd9\xed\x9a\xa2\x20\x64\xac\xa8\x24\x7f\x44\x1e\x03\xbe\x47\x8e\x43\x40\x57\xb7\xd2\x3c\xb4\xd6\xf5\xf9\x70\x90\x92\x5e\x1c\x30\x43\xc6\x6d\x09\x7a\x34\xe9\x31\x05\x7e\x9e\x4e\xb3\x09\xa6\x40\x0e\x5a\xf6\x17\x4c\x97\xcc\xea\x20\x50\x1c\x2f\x79\x6f\x07\x51\xbf\x6b\x48\xec\x4f\xb5\x28\x3f\xe0\xd4\x7c\xef\x46\xff\x53\xf4\x14\xc7\x4f\x9c\x2a\x07\x5c\xe1\xb6\xe6\x4f\x8b\xf9\x40\xcb\x58\x7b\xfe\x3b\x5c\x61\xd6\xb2\x79\x44\x97\x75\x30\xfd\x97\xa8\xef\x14\xe5\x1a\xe7\x6b\x00\x00\x00\xff\xff\x94\xa0\xd0\x5c\xb9\x03\x00\x00")

func ssmContainersRbacYamlBytes() ([]byte, error) {
	return bindataRead(
		_ssmContainersRbacYaml,
		"ssm-containers-rbac.yaml",
	)
}

func ssmContainersRbacYaml() (*asset, error) {
	bytes, err := ssmContainersRbacYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ssm-containers-rbac.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb4, 0xd, 0xb0, 0x65, 0xfa, 0x77, 0xe3, 0x24, 0x33, 0x29, 0xb9, 0xa7, 0xbf, 0x2e, 0x5, 0x97, 0x0, 0xe3, 0x91, 0x5, 0xd, 0x74, 0xde, 0x15, 0x27, 0xec, 0x71, 0xe5, 0x5d, 0x66, 0xf9, 0xa6}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"emr-containers-rbac.yaml": emrContainersRbacYaml,
	"ssm-containers-rbac.yaml": ssmContainersRbacYaml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"emr-containers-rbac.yaml": {emrContainersRbacYaml, map[string]*bintree{}},
	"ssm-containers-rbac.yaml": {ssmContainersRbacYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
