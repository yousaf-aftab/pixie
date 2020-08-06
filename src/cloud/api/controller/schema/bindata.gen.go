// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.graphql
package schema

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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x4f\x6f\xe3\xb8\x15\xbf\xeb\x53\xbc\x6c\x0e\x9b\x00\x89\x0f\x45\xbb\x07\x9f\xaa\xda\xde\x1d\x75\x12\x27\x8d\x9d\x4c\xb7\x8b\x20\xa0\xc5\x67\x8b\x88\x44\x6a\xc8\x27\x27\x6e\x31\xdf\xbd\x78\x24\x25\x4b\xb6\xd3\xc5\x2c\xd0\x93\x45\xf2\xfd\xf9\xf1\xfd\xa7\xcf\x61\x59\x28\x07\x6b\x55\x22\x48\x74\xb9\x55\x2b\x74\x40\x05\x82\xcb\x0b\xac\x04\xac\xad\xa9\xfc\x3a\xbd\xcf\xc0\xa1\xdd\xaa\x1c\x47\xc9\x79\x72\x0e\x19\xfd\xe8\x40\x1b\x02\x25\x51\x94\x57\xb0\x6a\x08\xde\x10\x34\xa2\x04\x32\x50\x09\xdd\x88\xb2\xdc\xc1\x06\x35\x5a\x41\x08\xb4\xab\xd1\xc1\xda\x58\x2f\x6f\xb9\xab\x71\x91\x5b\x55\x13\x3c\x66\xc9\x39\xbc\x15\xa8\x81\x3a\x30\xca\x41\x53\x4b\x41\x28\x47\x01\x62\x2e\x34\xac\x10\xa4\xd1\x08\xab\x1d\xd8\x46\x6b\xa5\x37\xe3\xe4\x1c\x60\x63\x45\x5d\x7c\x2d\xaf\x03\xe4\x6b\xaf\x27\x48\x6e\x75\x5f\x93\x8b\x17\x1a\x45\x62\xb8\xbe\x36\x0d\xd5\x0d\xb5\xfb\x72\x44\xce\xc3\x50\x79\x01\x6f\xaa\x2c\x7b\xc0\x0b\x84\x48\xcc\xb2\x03\x40\x2a\x04\x05\xba\x15\x42\xad\xf2\x57\x94\xd0\xd4\x0c\x8d\xc9\x1f\xb3\x51\x12\x6d\xdb\x93\xef\x39\x1d\xb8\xc2\x34\xa5\x04\x7c\x57\x8e\x40\xe9\x60\x6e\x51\x21\x48\x65\x31\x27\x63\x77\x20\xfa\x4e\xe8\x30\x33\xfb\x28\x49\xa2\x6b\xfe\x93\x00\x7c\x6d\xd0\xee\xc6\xf0\x0f\xfe\x49\x00\xaa\x86\x04\x29\xa3\xc7\x70\x1b\xbf\x92\x6f\x49\xe2\x41\x3f\x3a\xb4\x99\x5e\x1b\xcf\xa6\xe4\x18\xb2\xe9\x59\x02\xa0\x45\x85\x63\x58\x90\x55\x7a\xc3\x6b\xac\x84\x2a\xfb\x1b\xb5\xca\xa9\xb1\x03\x1a\x63\x37\xf3\x01\xdb\xb7\x24\x41\xdd\x54\x90\x5a\x52\x6b\x91\x13\xfb\xd6\xeb\x01\x48\x97\x2f\x8f\xf3\xcf\xf3\xbb\x2f\xf3\x76\x79\x93\xcd\x1f\xff\xf9\x92\xde\x4e\x7f\xfa\x73\xbb\x35\x4d\x1f\xbe\x64\xf3\xe1\xde\xe4\x6e\xbe\x4c\xb3\xf9\xec\xe1\x65\x31\x5b\xbe\xfc\x9a\xde\xde\x2c\x4e\x1f\xf5\xe5\x75\x40\x1a\x32\xb9\xa9\xea\x12\x09\xd3\x9c\xed\xd0\x41\x4a\x07\x88\xce\x21\xd5\x80\x52\x11\x08\x4f\x06\x26\xcf\x1b\xeb\x40\xad\x41\x40\xe3\xd0\x42\x21\x1c\x54\x46\xaa\xb5\xe2\xb8\x2e\x10\x94\xf6\x81\x80\xef\xc4\xce\x56\xda\xa1\x25\xa5\x37\x60\x2c\x48\x2c\xd1\x7f\xe7\x85\xb0\x22\x27\xb4\x6e\xe4\x95\xf8\x40\x50\x3a\x2f\x1b\xc9\xe9\xb5\xab\x3d\x43\xf0\xfc\x2b\xee\x56\x46\x58\x09\x42\x4b\xa8\x85\x0b\x02\x4c\x55\x09\x2d\x3d\x3b\x23\x9e\x4d\xb3\x65\x80\x0b\x0e\x4b\xcc\xf7\x78\x75\xb9\x3b\x0d\x3a\x2f\x8c\x43\x0d\x42\x83\xe8\x59\x03\x5c\xb3\xd9\xa0\x63\xde\x51\x0b\x4b\xaa\x5c\x10\xe3\x32\x5e\x05\x83\x1a\xb0\xf8\x50\x57\xd4\xc6\x6d\x65\xb6\x21\x27\x58\xd5\x8f\x0e\x58\x37\x27\xb5\xf1\x9b\x9a\x0d\x23\xea\xda\x9a\xda\x2a\x9f\x3d\x62\xd5\xde\x62\x31\xbb\x99\x4d\x96\x27\xbd\x34\xd3\xa4\x68\xf7\x59\x69\x19\xbc\x34\xfb\xdc\xf3\x12\xaf\xee\xef\xa6\xf1\x6b\xf1\x34\x69\xbf\x26\x0f\xd9\xfd\x32\x2e\xe6\xe9\xed\x6c\x71\x9f\x4e\x66\x5d\xc8\x4f\xb1\x2e\xcd\xae\x42\x4d\x9f\x71\x77\x10\xf7\xaf\xb8\xeb\x87\x74\x6e\x91\x6b\x4d\x4a\xb7\x6e\x0c\x3f\x97\x46\x10\xef\x72\x45\x1c\x44\xb9\x17\xeb\x93\xcd\x8b\x63\x03\x8c\xbb\xcc\x3a\x0b\xae\xbe\x9b\xde\x5d\x70\x50\x58\xa5\xcd\xe5\x18\x6e\xc5\x2b\x42\x36\x05\x8b\x5f\x1b\x65\x51\x82\xd1\x39\xd7\x07\x6f\x46\x07\x66\x8b\xde\x74\x55\x53\x92\xba\xce\xcb\xc6\x11\x5a\x70\x4d\x5d\x1b\x4b\x6c\xb7\xb8\x75\x11\xa0\x5f\x8e\x61\x12\x36\x5a\x8d\xf1\xdc\x8d\xe1\xb7\xfe\xc9\xf3\xff\x15\xcd\xc4\x68\x8d\x3e\x00\x8f\x70\xed\x8f\xf6\x08\x55\x5b\x18\x2e\x44\xaf\x42\x8c\x07\xf5\x82\x25\xdc\x64\xed\x0e\xf3\xb5\xb4\xae\xe3\xea\x57\x9d\xcb\x3d\xbb\x6b\x35\xf5\xa3\xf6\xc2\xe7\x69\x4b\x7d\x15\xa3\xf4\xde\xb8\x31\x64\x9a\xae\x62\xfe\x8c\x3f\x28\x15\x57\xed\x4d\x1f\xb3\x69\x5f\x63\x8f\xf8\x01\x5d\x53\xd2\xa1\xda\x9f\x15\x96\xf2\x50\xf7\x9a\x37\xe3\x95\x4f\xc6\xfc\x95\x2f\x6d\xad\x53\x52\xbb\x61\x62\x76\xe9\x69\xf2\xe7\xd3\xf0\x06\xd4\x8b\x2e\xcf\x9f\x13\x1f\x0a\xa1\xd9\x56\x1b\x0b\xa8\x65\x6d\x94\x26\x77\x05\x16\xd7\xc1\xe3\xd2\xe4\x5c\x0a\x20\x2f\x4d\x23\x45\xad\x46\xb5\x35\xbe\x1e\x94\x6a\x8b\x4f\x0a\xdf\x18\xcd\x4d\xfc\xbe\x45\x12\x52\x90\x08\x51\xd6\x52\x4c\x8c\x26\xd4\xe4\x62\x48\x9c\x5d\x8e\xe1\xe6\xe0\x88\xc9\x43\x6b\x66\x71\x01\xd1\x50\x58\x38\x3d\x21\x6a\x31\x38\x38\x0b\x77\x0a\x19\xce\xc9\xec\x7c\xba\xf6\xf2\x9d\x15\x0c\x0a\x40\x90\x3f\xa0\xe9\x89\x1f\x92\x76\xb9\x7e\xec\x70\x9f\xf8\x5c\xd1\x91\x47\x99\x4a\x10\xa1\x8c\x3d\x41\xb9\x5e\x83\x70\xd1\xf7\x61\xa0\xe0\x82\xbc\x42\xd4\x50\x0b\xeb\x50\xb6\x63\xc2\xb0\xcc\x9a\xae\x16\x87\x3a\x2c\x56\x0b\x32\x35\xd4\xc6\x29\xf6\xa3\x6f\x06\x9d\xce\xac\x1f\x62\x9e\xfe\x4b\x81\x54\xa0\x3d\xc2\xc0\xb8\x04\x6c\x45\xa9\xe4\x15\xe0\x3b\xe6\x0d\x89\x55\x89\x6d\x8f\x61\xa9\xca\xcd\xba\xfd\x31\xfc\xcd\x98\x12\x85\x0e\xfd\xa6\x2c\x7b\x2d\x23\x8c\x6f\x28\xf2\x02\xcc\xda\x2b\x8a\x20\x3d\x36\xfe\xde\x93\x8e\xe1\xb7\x65\x7f\xe3\xb9\x33\xea\x60\xbb\x67\x4f\xa5\x25\xbe\xf7\x04\x87\xc6\x43\x05\x3a\x1c\x60\x10\xd6\xdb\x3e\xaa\xcc\x98\xcb\x27\xf5\xc0\x0a\xa1\x4d\xf2\xf5\x45\x8f\x39\x8e\x9f\xec\x29\xb1\x8a\x0a\xfd\x10\x57\x71\x61\x64\xbd\xd1\x2a\x3d\x43\xb1\x9e\xfd\x2a\x5d\x13\xda\x85\x17\xde\xb7\x94\x1b\x5c\xfc\xa3\x44\x3c\x15\x56\x07\xa6\x78\x55\x5a\x7e\x54\x26\x0e\xe6\xb5\xd8\xa1\x38\x2f\x7c\x29\xeb\x76\x2b\x41\x79\xc1\x21\x22\xf1\xdd\x97\x91\x4c\xd3\x73\xd7\x76\x63\xa1\x5e\x90\xa0\xc6\xf5\xcc\x2f\x71\x2d\x38\xc0\x1d\x71\xdb\x56\x6b\x1e\xee\x8b\x18\x3f\xaf\xda\xbc\x69\x36\xc4\xd3\xbf\x5e\x16\xc3\x01\x8a\x59\x23\x8b\x83\x02\x45\x49\xc5\x8e\xb9\x0b\x14\x96\x56\x28\x28\x38\xcc\x62\x8e\x6a\xeb\x3b\x0e\x58\xdc\x34\xa5\xb0\xa0\x34\xa1\xdd\x8a\xd2\xf9\xd9\x87\x8a\x10\xf7\x6d\xdb\x51\x0e\x2c\xba\xda\x68\xc9\x20\xc8\xf8\xfa\x88\x8e\xdc\x1e\xc7\xa7\x59\x7a\xb3\xfc\xf4\xeb\x01\x8e\x30\xbd\x1b\x5f\xd6\x94\xcb\x43\x43\xe2\x2c\x0d\x91\xf5\xcb\xc3\xfd\x04\xf2\xae\x4d\xc1\xca\xa2\x78\x75\x23\x2f\xa0\x30\x35\x86\x3c\x16\xd4\x0d\x43\x2d\x20\x2f\x37\x37\x15\xc2\x4a\xe4\xaf\x3c\x7a\x29\x8d\x1e\xba\x45\xd7\x54\x1c\xc0\x10\x11\x05\x24\x7b\xa0\xd3\x6c\x31\xb9\x9b\xcf\x67\x93\xe5\x6c\x7a\x6c\x35\xff\xd2\xe1\x4b\xc6\x47\x10\xf6\x6d\x10\x1f\x08\xb5\x35\x39\x3a\xc7\xe9\xd1\x92\xf7\xfc\x71\x3f\x4d\x97\xd9\xfc\x97\x4e\xf4\x56\xfd\x5b\xb5\x73\x60\x7b\xff\xf0\x44\xe3\x2d\x7e\xb5\x39\xd4\x04\x42\xef\xc0\xf8\x74\x59\x37\x36\xa4\x4d\x88\x8a\xf0\xf6\x72\x20\x56\xa6\x09\x86\x78\x8b\x79\xa5\xa8\xef\x67\x63\x4f\xa0\x39\xbe\x69\x84\xf3\xc6\xcf\x1a\xbb\x8b\xee\x0c\x3a\x02\xaa\xb5\x50\x25\x86\x11\x58\x31\xbe\x37\xbe\xb6\x80\x95\x90\x87\x96\xf4\x57\x9d\xbd\xfc\x9c\x66\x37\xb3\x69\x97\x50\x4f\x5e\xc1\xc4\xe8\xb5\xda\xf8\x90\xae\x85\x73\x54\x58\xd3\x6c\x8a\x99\xe6\xbc\x95\xfb\x6c\x6d\x99\xb8\x99\x08\xa5\x07\xa9\x70\xf8\x1e\x3a\x3d\x18\x7a\x4c\x7d\xb2\x0a\x9d\x13\x9b\x7e\x66\x5a\x14\xae\x97\x94\xad\xce\x7b\x23\xff\x90\xb6\xc6\x7d\x97\x3a\x60\xb7\x87\xcb\xf9\xc9\x70\x78\x53\x6e\x85\x9d\x11\xf6\x43\xe3\xc1\x88\xdc\xaa\x1d\xd4\x0b\xdf\xf0\x85\xa3\x4f\x6d\x76\x0f\x90\x6e\x7b\x6e\x18\x0f\x9c\xb2\x3f\x7d\x42\xeb\x86\xe5\x2a\x06\xfb\x87\x07\xf3\x61\xcd\xab\x2d\x12\xed\x26\x27\xcf\x8e\xc7\xa2\x68\x0a\x6b\xca\xfb\x52\x68\xec\xec\xef\xeb\x62\xb7\x0a\xc3\x81\x6e\xaa\xb9\x91\x18\x46\xc4\xb8\x91\x69\x47\xb6\xe1\xc1\x00\x65\xff\xf0\xc0\x7e\xc3\xb1\x37\x58\xb2\x4e\xa5\xb4\xe8\x06\x9e\x23\xf3\x8a\xfa\xf8\x4d\xd1\x3e\xd7\x3d\xe3\xc4\x47\x41\x14\x3c\x18\xf8\xe1\xaf\x12\x6b\x8b\x3c\x22\xc8\x8b\xd6\xe5\x3f\x44\x82\x50\x67\x39\x77\x62\x18\xc1\x56\x09\xa8\xdf\xe3\xcc\xf3\xc3\x65\x02\xf0\xe8\x73\xae\xef\x98\x8b\x68\x32\xb6\x58\x36\x3d\xbb\xfa\x5f\x99\x73\xd9\x7d\x9d\x75\x30\x07\x83\xd3\xd1\x1c\x05\x30\xe5\x37\xf1\x90\xaa\x37\x76\x75\xe2\x3a\x73\xee\xdf\x00\xf1\x1f\x84\xc6\x0e\xfe\x8f\x00\x70\x85\xf8\xd3\x5f\x7e\x3a\xb6\xe1\xe0\x39\x10\x3c\x40\x58\xf9\x76\x1c\x4f\x9e\x8f\x68\x3d\xd9\x76\x18\x78\x3e\x11\x0b\xa1\x37\x58\x9a\xcd\xc0\x77\xaa\x42\x47\xa2\xaa\x7b\x31\xff\x2d\x49\xce\xe1\xe1\x77\xa6\x68\xaf\xf2\x70\x78\xfe\x9d\x3f\x62\x8e\xde\x9e\xdf\xa9\xa6\x9d\x94\xbd\x9a\x2a\xea\x1c\x1f\xa1\xf0\x7f\xf1\xbc\x97\x2d\x75\x1f\xc1\x56\xb9\xbf\x2f\xee\xe6\x7f\x04\xc4\x70\xb2\xff\xae\x9b\x02\x77\xa8\x16\xe5\x30\x40\xbe\x4b\xf9\x07\xf7\x3f\x78\x73\xc4\xea\x30\xbc\xfa\xb7\xe4\xbf\x01\x00\x00\xff\xff\xb7\xe2\x0c\xbe\x17\x15\x00\x00")

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

	info := bindataFileInfo{name: "schema.graphql", size: 5399, mode: os.FileMode(436), modTime: time.Unix(1596744439, 0)}
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
