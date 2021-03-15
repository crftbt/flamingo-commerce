// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package graphql generated by go-bindata.// sources:
// schema.graphql
package graphql

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x98\x5c\xb6\x01\x45\x77\xf7\x6d\x6d\xb6\x21\x40\x51\x6c\xed\x6e\xc3\x50\x28\x16\x6d\x13\x93\x29\x43\xa2\xd7\x06\x43\xff\x7d\xb0\x24\x27\x76\x63\x77\xc1\x96\x53\x4c\x3e\xbf\x47\xd1\x8f\x94\xec\x5b\x84\x6b\xdb\x34\xe8\x0a\x7c\xb8\xee\xbc\xd8\x06\xdd\xc3\xbd\x28\xe9\xfc\xc3\x1d\xfa\xce\x08\xfc\xce\x00\x00\xc8\xdf\xd8\xaa\x42\xbd\xe5\x1c\xae\xac\x35\xa8\x78\x15\x12\x9d\x47\xb7\xdd\xe4\x70\x2f\x8e\xb8\x5a\x65\xcf\x59\xb6\x40\x3b\xe5\xd3\xc7\x57\xfa\xe7\xf5\x00\xf3\xd0\xa2\xf3\x96\x95\x01\xad\x44\xad\x43\x76\x08\x6d\x94\xa8\x7c\x86\xfa\x4b\xc8\xf7\xd9\xc4\xf6\x19\x05\x14\xf8\x16\x0b\x2a\xa9\x00\xa5\xb5\x43\xef\xa1\x74\xb6\x01\xa9\x11\x8a\xf4\x66\xa4\xaf\x50\x3e\x44\xc4\xdb\xbe\xae\xed\x66\xf5\x6e\x4e\x25\x61\xa2\x42\x7a\x40\x0f\x52\x2b\x99\x90\x42\xeb\xec\x2f\xd2\xa8\x2f\xa0\x50\x0c\x3b\xec\x9b\xa4\xa1\xb4\x0e\x76\x64\x0c\x71\x05\xef\xc1\xd7\xd4\xb6\xc4\x55\xac\x40\x0d\x6c\x39\x7c\x5f\x14\x5e\xfd\x88\xd2\xdf\x6a\x04\x8d\xa5\xea\xbb\x39\xd0\x1c\x8e\x68\xcb\x49\x2d\x17\xc0\x9d\x31\x40\x21\xea\x10\xc8\x03\x5b\xc6\xa8\x9a\x48\xee\x13\x47\x92\xf9\xeb\xd1\xc7\xfa\xc3\x81\xfe\x43\xfe\x2a\x52\x9c\xa1\xbe\xec\xad\xa3\x01\x92\xbf\x2a\x64\x8d\x2e\xef\xff\x4e\x7c\x56\x92\xf3\x72\xab\x1a\xcc\xa7\x71\xa3\x0e\xe1\x49\xbc\x21\xad\x0d\xc6\xcc\x24\xae\x88\x3f\x36\x8a\xcc\x0b\x9e\xd6\x61\x49\x4f\x51\x77\x92\xd8\x91\x93\x5a\xab\x7d\x48\x6d\x94\x60\x88\xb2\x12\xea\x8d\x4d\xb2\x3f\x67\x84\x52\x23\x46\x33\xb4\xdd\xac\x06\x03\x51\xa4\x4a\x98\x1b\xe2\xe0\xa6\x44\x1a\xad\x53\x04\x9d\xd3\xdf\xb8\xd0\xc2\x36\xad\xe2\x53\xd8\x14\xd3\xb1\xb8\xfd\xb5\xd5\x98\x2f\x61\xd6\x9f\x8c\xaa\xe2\x97\x27\x7f\xf0\x87\xaf\x6d\x67\xf4\x61\x26\x94\x0f\x76\x59\x70\xd3\x9c\x4d\x8e\x7a\x93\x35\xf4\x2f\x6a\x2f\x67\x67\x76\x28\xf2\x59\xb9\xb1\x8d\x96\x1a\x30\xb6\xd4\x12\xa6\xb5\x5e\x4e\xba\xf8\x12\x33\xb6\xd4\x02\xc6\x61\x45\x96\x4f\x99\xc6\x18\x2f\x0e\x51\x5e\xe7\x89\x98\xdb\xae\xd9\x0d\xc3\x73\x8a\x11\x34\xd8\xd6\x96\x5f\x3b\x3b\xa6\xd1\x58\xd2\x7a\xce\x32\x7c\x12\x64\x0d\xc1\xe9\x5f\x3b\x74\xfb\xe4\xea\xf5\x3a\x7e\x87\x3b\x94\xce\x71\xfc\x62\x26\xdc\x3d\x40\x0c\x3e\xdc\x4d\x61\x97\xc6\x3d\xe3\x1c\xb2\x80\x47\xef\xc9\xf2\x84\x60\xe9\x62\x9b\xdb\x30\x93\x2b\xef\x8c\x32\x0e\xab\x7e\xa1\x10\xb0\x0e\x14\x03\x3a\x67\x5d\xef\x4a\x92\xb8\xfa\xe4\xc8\x71\x19\xe8\xb7\x25\xec\x6d\x07\xda\xf2\x1b\x81\x47\xc5\x02\x62\xa1\x56\xac\x0d\x06\xde\xc0\x70\x01\x45\x8d\xc5\x4f\x78\x24\xa9\x17\x8b\x8f\x9e\xbc\x7c\xbd\x05\x73\x67\x4f\x87\x7e\xce\xfe\x04\x00\x00\xff\xff\xf2\xfb\x6b\xd8\x0f\x08\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
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
	"schema.graphql": schemaGraphql,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}