// Code generated by go-bindata.
// sources:
// templates/call.tmpl
// templates/function.tmpl
// templates/header.tmpl
// templates/inline.tmpl
// templates/inputs.tmpl
// templates/message.tmpl
// templates/results.tmpl
// DO NOT EDIT!

package render

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

var _templatesCallTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8d\x41\x0a\xc2\x40\x0c\x45\xaf\x12\x4a\x17\x0a\x25\x07\x10\x3c\x80\x1b\x11\x15\xf7\x61\x9a\x6a\xa0\x8e\x92\x46\x45\x42\xee\x6e\x3b\x54\xba\x0a\xfc\xff\xf3\x9e\x7b\xcb\x9d\x64\x86\x2a\x51\xdf\x57\x11\xee\x1f\xb1\x1b\xe0\x91\x13\xcb\x9b\x75\x4a\xa4\x83\xfc\x30\xc0\xdd\x70\x32\x7d\x25\x8b\x30\x43\x77\xce\xed\xd4\xfe\x97\x80\x11\x4b\x8a\x7b\xba\x73\xc4\xca\x5d\x29\x5f\x19\x6a\x69\xa0\xe6\x1e\x36\x5b\xc0\x03\xe9\x58\x1a\xeb\x30\xd3\x6b\x89\x68\x60\xfe\x2d\xec\x32\x99\x88\xa5\xc7\xf3\xf7\xc9\xa3\xfe\x42\x2a\xd4\x4a\x1a\x45\xb8\xa8\xca\x59\xcf\xf7\x17\x00\x00\xff\xff\x2d\x6e\x46\xc9\xd1\x00\x00\x00")

func templatesCallTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCallTmpl,
		"templates/call.tmpl",
	)
}

func templatesCallTmpl() (*asset, error) {
	bytes, err := templatesCallTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/call.tmpl", size: 209, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x4d\x6f\xdb\x3c\x0c\x3e\x3b\xbf\x82\x0d\xda\x17\xf1\x0b\x43\xbd\x67\xc8\x61\x43\xbb\xa1\x87\x2d\x45\x16\x6c\x87\x61\x18\x34\x9b\xce\x04\x28\xb2\x27\xc9\x29\x0a\xc3\xff\x7d\xd4\x87\x3f\x9a\x26\xd9\x4e\xb6\x28\xf2\xe1\xc3\x87\xa4\xda\xb6\xc0\x52\x28\x84\x79\xd9\xa8\xdc\x8a\x4a\xcd\xbb\x0e\xda\xf6\xba\x84\xe5\x0a\x58\xd7\xcd\x66\xee\x82\x2c\x6c\x8b\xc6\x7e\xe2\x7b\xec\xba\x85\x85\xff\x2d\x9d\x84\xda\xb1\x6d\x0a\xed\x2c\x71\x27\xe3\x22\xbe\x7d\x37\x56\x37\xb9\x85\x96\x42\x44\x09\xaa\xb2\x70\x5d\xb2\x47\x2d\x94\x7d\x50\x75\x63\x0d\x41\x26\x89\x22\x1c\x20\x4f\x42\x20\x3f\x54\x85\x4f\xfa\x24\xec\x2f\x60\x1b\xcc\x51\x1c\x50\x7b\x13\x41\xb0\x07\xf3\xd9\x63\x7a\x83\xe6\x6a\x87\xc0\xde\x0b\x94\x45\xc0\x22\x6a\x81\x96\x27\xf9\x5c\x87\xbf\x1e\x13\xa5\xc1\xe8\xd6\x03\xbb\xb2\x4e\xfb\xc6\x6f\xcc\xf1\xc8\x35\xe1\x5a\xd4\x7d\x1e\x6f\x38\x1b\x1d\xa3\x36\x68\x1a\x69\xfb\x90\xaf\x5c\xd9\xb3\x11\xae\xb8\x0d\xda\x46\x2b\x73\xaf\x75\xa5\x7d\xcc\x13\x45\xd0\x09\x7e\x56\x95\xec\x7d\x67\x49\x47\x22\x27\xb7\xb7\xb0\x5d\xdf\xad\x97\xf0\xb6\x28\xc0\x49\x0e\x39\x37\x68\x18\x5d\xcf\x92\xb2\xd2\xf0\x23\x03\x6b\x5d\x1b\x02\x97\xd0\x95\xf6\x84\xb2\xc7\xc2\xbe\xd6\x87\x40\x82\x93\x63\xed\x3d\x39\xc5\xfd\x17\x09\xc5\x6a\xd8\x17\x2e\x1b\xaa\xa9\xbd\xdc\x98\x25\xb1\x62\xc3\x29\x1b\xab\x4a\x4e\xa9\x4f\x49\xb9\x2a\x60\xe1\x46\x87\xad\x95\x7c\x9e\x4a\x94\xbe\xb6\xaf\x15\x7a\x1a\x29\xc4\xac\x16\xf7\xb5\xe4\x96\x66\x5a\x87\x5e\xcc\x69\x04\x3d\xf4\x78\x93\x73\x29\x07\xf3\xc5\x7e\x90\x31\xdc\x1c\x53\xa1\x10\xa4\x3e\x79\xa1\x4e\x01\xbf\x19\x90\x17\xce\xef\x6a\x05\x4a\xc8\xd4\x7d\x49\x8d\xbe\xcb\xae\xad\x89\x65\x1e\xb0\x5c\xcc\xa7\x48\x7b\x34\x86\xef\x30\xb2\x44\xe7\x01\x2b\xb8\x39\x64\xd0\x07\xdf\x1c\xe6\xd9\x8b\xe4\xc2\x6f\xd8\x18\x91\x4d\x52\xa5\x43\x81\xe3\x7c\x26\x79\xa5\x68\x8b\x1b\x3c\xd3\x92\xe3\x91\x1e\x16\xf2\x1d\x37\x22\x0f\xf3\xdc\x4b\x34\x65\x21\xe9\x45\x19\xd4\xfd\x50\x85\x0d\x08\x85\x4f\x57\x62\xba\x9e\x97\x30\xae\x34\x96\x12\x73\xcb\xee\x10\xeb\xfb\xdf\x0d\x97\x8b\x01\x36\x7b\x09\x9a\x06\xd4\x58\xcb\x3f\x29\xeb\x6b\xa2\x37\x2a\xf6\xf6\x23\x95\x2a\x6a\xe9\xa6\x7a\xa0\x1e\x11\x47\xf5\xff\x22\xfd\x59\x76\x53\x7d\xdd\xd6\xd2\xf3\xda\x9f\xfe\x04\x00\x00\xff\xff\xa0\x60\x44\xbd\x8a\x05\x00\x00")

func templatesFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesFunctionTmpl,
		"templates/function.tmpl",
	)
}

func templatesFunctionTmpl() (*asset, error) {
	bytes, err := templatesFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/function.tmpl", size: 1418, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\xaa\xad\xe5\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xa8\xae\xd6\x0b\x80\x30\x81\x82\x5c\x99\xb9\x05\xf9\x45\x25\x0a\x1a\x5c\xd5\xd5\x45\x89\x79\x40\x69\x3d\x4f\xb0\x48\x71\x6d\x2d\x50\xa1\x5f\x62\x2e\x50\x15\x44\x4b\x49\x06\x50\x7d\x75\x75\x6a\x5e\x0a\x90\xd6\x84\xb3\x00\x01\x00\x00\xff\xff\x81\x22\x53\x6f\x6b\x00\x00\x00")

func templatesHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeaderTmpl,
		"templates/header.tmpl",
	)
}

func templatesHeaderTmpl() (*asset, error) {
	bytes, err := templatesHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/header.tmpl", size: 107, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\xcb\x01\xd2\x4a\xb5\xb5\xd5\xd5\x99\x69\x0a\x7a\xfe\x79\x39\x95\x41\xa9\x25\xa5\x45\x79\xc5\xfe\x79\xa9\x61\x89\x39\xa5\xa9\xb5\xb5\xe9\xf9\x25\x0a\x56\xb6\x0a\xd5\xd5\x25\xa9\xb9\x05\x39\x89\x25\x40\x6d\xc9\x89\x39\x39\x4a\x0a\x7a\xb5\xb5\xd6\x40\xe1\xd4\xbc\x14\x90\x7e\x30\x05\x08\x00\x00\xff\xff\xfc\x48\x08\x64\x5a\x00\x00\x00")

func templatesInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInlineTmpl,
		"templates/inline.tmpl",
	)
}

func templatesInlineTmpl() (*asset, error) {
	bytes, err := templatesInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inline.tmpl", size: 90, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInputsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcb\x41\x0a\x02\x51\x08\xc6\xf1\xab\xc8\x30\xcb\x78\x07\x08\x3a\x40\xbb\xae\xf0\x60\x3e\x43\x28\x09\xc7\x56\x1f\x73\xf7\x46\x57\xad\x94\x9f\x7f\xc9\x0d\x6a\x0e\x59\xcc\x3f\xdf\xdc\x97\xe3\x20\x57\x95\xeb\x4d\x46\xad\xa6\xb2\xea\x78\x84\x79\xde\x3b\x28\x8c\xe9\x4f\xb4\xcf\x98\x6f\x24\xe2\xe4\xcc\x41\x36\xd4\xe7\x45\x48\xf8\x56\x35\x5e\x3b\xfa\xec\x67\xfb\xef\x35\x7e\x01\x00\x00\xff\xff\x43\x89\x5c\xae\x80\x00\x00\x00")

func templatesInputsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInputsTmpl,
		"templates/inputs.tmpl",
	)
}

func templatesInputsTmpl() (*asset, error) {
	bytes, err := templatesInputsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inputs.tmpl", size: 128, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMessageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xd1\x0a\x82\x40\x10\x45\x7f\x65\x10\x85\x02\x99\x0f\x08\xfa\x80\x5e\x22\x22\x7a\x5f\xf2\x6a\x03\xba\xc9\xee\x6a\xc4\x30\xff\x9e\x8a\x09\x3d\x5d\x38\x1c\xce\x8c\x6a\x85\x5a\x3c\x28\xeb\x10\xa3\x6b\x90\x99\xa9\x4a\x4d\xfe\x95\x88\x2f\x41\x7c\x3a\xf9\x7e\x48\xd1\xac\x18\x99\x54\xe1\xab\xd9\x78\x4b\x7a\x12\x5f\xf1\x80\x8c\x08\x33\xe1\xdb\xa7\x07\xdf\x5d\x3b\xc0\x8c\x37\x91\xcf\xae\x9b\xc0\x6e\x89\xfe\x07\x55\x83\xf3\x0d\x28\x97\x92\x72\xb4\x74\x38\x4e\x82\x0b\x93\x9f\x10\xe2\xfa\x47\x2e\x66\xe5\xef\x6e\x31\x6e\xdd\x65\xf6\xeb\x7e\x03\x00\x00\xff\xff\x77\xb1\x17\x14\xc6\x00\x00\x00")

func templatesMessageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMessageTmpl,
		"templates/message.tmpl",
	)
}

func templatesMessageTmpl() (*asset, error) {
	bytes, err := templatesMessageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/message.tmpl", size: 198, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResultsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcd\x41\x0a\xc2\x40\x0c\x05\xd0\xab\x7c\x4a\x97\xa5\x07\x10\x5c\x8a\x7b\x6f\x20\x34\x95\x81\x92\x81\x3f\xd3\x55\xc8\xdd\x4d\x6a\x51\x70\x35\x93\xfc\x97\xc4\x6c\x91\xb5\xa8\x60\xa0\xb4\x7d\xeb\x6d\x70\x87\x19\x9f\xfa\x12\x8c\x65\xc2\x28\x1b\x2e\x57\xcc\x8f\x4f\xec\x6e\x56\xd6\x48\xdc\xa7\x70\xa2\x4b\x76\xee\xb5\x63\xce\xcf\x59\x87\x88\x81\xbe\x53\xdb\x8d\xac\x4c\x2c\xe4\x99\xe3\x00\x95\xdf\xa5\xff\x38\x0f\xfe\xec\xf1\xbe\x03\x00\x00\xff\xff\xb0\x4f\xcf\x61\xa8\x00\x00\x00")

func templatesResultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesResultsTmpl,
		"templates/results.tmpl",
	)
}

func templatesResultsTmpl() (*asset, error) {
	bytes, err := templatesResultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/results.tmpl", size: 168, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
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
	"templates/call.tmpl": templatesCallTmpl,
	"templates/function.tmpl": templatesFunctionTmpl,
	"templates/header.tmpl": templatesHeaderTmpl,
	"templates/inline.tmpl": templatesInlineTmpl,
	"templates/inputs.tmpl": templatesInputsTmpl,
	"templates/message.tmpl": templatesMessageTmpl,
	"templates/results.tmpl": templatesResultsTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"call.tmpl": &bintree{templatesCallTmpl, map[string]*bintree{}},
		"function.tmpl": &bintree{templatesFunctionTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{templatesHeaderTmpl, map[string]*bintree{}},
		"inline.tmpl": &bintree{templatesInlineTmpl, map[string]*bintree{}},
		"inputs.tmpl": &bintree{templatesInputsTmpl, map[string]*bintree{}},
		"message.tmpl": &bintree{templatesMessageTmpl, map[string]*bintree{}},
		"results.tmpl": &bintree{templatesResultsTmpl, map[string]*bintree{}},
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

