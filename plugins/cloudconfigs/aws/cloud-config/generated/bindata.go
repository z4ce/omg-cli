// Code generated by go-bindata.
// sources:
// files/disk_types.yml
// files/vm_types.yml
// DO NOT EDIT!

package generated

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

var _filesDisk_typesYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x93\xbd\x6e\x83\x30\x14\x46\x77\x9e\xc2\x1b\x53\x25\xfb\xe2\xbf\xf2\x32\xa8\x2a\xa8\x42\x4d\x82\x05\x64\x48\x9e\x3e\xb6\x83\x32\xf8\x7a\xc8\x97\xf5\x08\x1d\x71\xbf\x03\x5f\xe2\xf2\x73\x9e\x7a\xd1\x2a\x49\xba\x6d\x84\x18\xe7\xed\x7f\xd8\xe6\x7b\x64\x09\x45\xf2\x7b\x5a\xae\xe3\x10\xd6\x25\x4c\xeb\x3e\x4f\x5b\x1f\x99\x10\xfb\x2d\xc4\x47\xfe\x02\x35\x2f\x05\x49\xed\x0b\x45\x42\x88\xc2\x28\x92\x85\x22\x21\x44\x91\xde\xba\x74\x64\x86\x9e\x52\x4a\x32\x43\x24\x9d\x74\xec\x9a\xcc\xd0\x45\x6a\x93\x40\x12\x67\x3d\x93\x64\x06\x0f\x5b\x5d\x16\xd3\x98\xce\x72\x4d\x86\x70\xa1\x6a\x22\xbc\x51\x35\x12\x5e\xa9\x9a\x09\xef\x54\x0d\x05\x96\xd2\xde\x38\xcb\x52\x65\x8a\x8d\xfc\xed\x94\x21\xb6\x72\xa6\xd0\x3e\xa4\xc9\xb3\x3f\xea\xa0\x1f\x9c\xc6\x3f\xc3\x27\x86\x54\xd6\x39\x47\x8a\xcd\x74\xe0\x37\x54\x8f\x00\x00\x00\xff\xff\x88\x86\x63\x21\x3d\x05\x00\x00")

func filesDisk_typesYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesDisk_typesYml,
		"files/disk_types.yml",
	)
}

func filesDisk_typesYml() (*asset, error) {
	bytes, err := filesDisk_typesYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/disk_types.yml", size: 1341, mode: os.FileMode(420), modTime: time.Unix(1474610230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesVm_typesYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x94\xd1\x6e\x83\x20\x14\x86\xef\xfb\x14\xbc\xc0\xc8\x04\x6b\x5d\x5f\xc6\x10\x3d\x73\x64\xa0\x04\x30\x73\x7b\xfa\xd1\x5a\x7b\xcd\xa1\x81\x78\x63\x02\xfe\xc0\x97\x73\x3e\x78\x23\x93\xd0\x70\x0d\xdf\x69\x3e\x11\xd2\xab\x79\x19\x3a\x63\x67\x03\xd6\x4b\x70\xd7\x30\x47\x88\x9c\x9c\x17\x53\x0f\x9d\xff\x35\x21\xeb\x19\x7d\xc4\x09\x01\xf3\x05\x1a\xac\x50\xdd\x20\xdd\xf7\x16\x27\x64\xcb\x8d\x86\x3d\xc6\x4e\xfe\x85\x71\x5b\x7d\x6c\x13\x56\xfc\x74\xcf\x3d\x9d\x9f\xad\x18\xc3\xef\x4f\xa1\x1c\x9c\x76\x20\x2d\x7b\x8b\x21\xda\xf3\xb9\x91\xa8\x15\xfa\x58\x58\x4e\x0b\xa5\x10\x48\x7b\x3e\x37\x12\xbd\x6d\x9a\x9d\xab\x6a\x78\x5b\xc7\xb7\x10\x06\xb9\x44\xf7\x4f\x73\xfa\x5c\x90\xd3\xab\xfb\x19\xe1\xa8\x83\x82\x61\xfa\x98\x4e\xc6\xd9\xa5\x69\xb1\x68\xbd\x59\x0e\x56\x33\x25\xec\x08\xd1\x4c\x35\x5d\xf7\x05\x39\xb5\xbf\x9f\x81\x12\xac\x28\x18\x4a\xb0\x9a\xa6\x81\x35\xe7\x33\x6f\x90\x60\x08\xbd\xfa\x42\x15\x5b\x0b\x09\x86\xbb\x8e\x6b\x8a\x61\xac\x24\x1a\xd2\xb1\xd4\x5e\xf2\xea\xfd\x12\xff\x58\xac\x29\x9a\x95\x29\x1b\xc3\x89\xf6\x02\x17\xee\x62\xb2\x04\xd5\x78\x61\x36\x8c\x6b\x2f\x14\x8e\x35\xac\xaa\xe3\x5f\x0e\x96\x64\x5b\x9d\xb5\x72\xff\x01\x00\x00\xff\xff\x64\x68\xbb\xcc\x7f\x0c\x00\x00")

func filesVm_typesYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesVm_typesYml,
		"files/vm_types.yml",
	)
}

func filesVm_typesYml() (*asset, error) {
	bytes, err := filesVm_typesYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/vm_types.yml", size: 3199, mode: os.FileMode(420), modTime: time.Unix(1478733001, 0)}
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
	"files/disk_types.yml": filesDisk_typesYml,
	"files/vm_types.yml": filesVm_typesYml,
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
	"files": &bintree{nil, map[string]*bintree{
		"disk_types.yml": &bintree{filesDisk_typesYml, map[string]*bintree{}},
		"vm_types.yml": &bintree{filesVm_typesYml, map[string]*bintree{}},
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
