package test

import (
	"os"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
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

func (fi bindataFileInfo) Name() string { _ = "STUB: not implemented"; return "" }

func (fi bindataFileInfo) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (fi bindataFileInfo) Mode() os.FileMode { _ = "STUB: not implemented"; return *new(os.FileMode) }

func (fi bindataFileInfo) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (fi bindataFileInfo) IsDir() bool { _ = "STUB: not implemented"; return false }

func (fi bindataFileInfo) Sys() interface{} { _ = "STUB: not implemented"; return nil }

var _viewsBlocksBlockTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xca\xc9\x4f\xce\x56\xaa\xad\xe5\xb2\xc9\x30\xb4\xf3\x48\xcd\xc9\xc9\xd7\x51\x00\x8b\x15\x2b\xda\xe8\x67\x18\xda\x71\x55\x57\xa7\xe6\xa5\xd4\xd6\x02\x02\x00\x00\xff\xff\xfd\xa1\x7a\xf6\x32\x00\x00\x00")

func viewsBlocksBlockTplBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func viewsBlocksBlockTpl() (*asset, error) { _ = "STUB: not implemented"; return nil, nil }

var _viewsHeaderTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\xaa\xad\xe5\xb2\xc9\x30\xb4\xf3\x48\xcd\xc9\xc9\xd7\x51\x48\x2c\x2e\x49\xac\xc8\x4c\x55\xb4\xd1\xcf\x30\xb4\xe3\xaa\xae\x4e\xcd\x4b\xa9\xad\x05\x04\x00\x00\xff\xff\xe4\x12\x47\x01\x34\x00\x00\x00")

func viewsHeaderTplBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func viewsHeaderTpl() (*asset, error) { _ = "STUB: not implemented"; return nil, nil }

var _viewsIndexTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xbd\x8a\xc3\x30\x10\x84\x6b\xeb\x29\xe6\xfc\x00\x16\xb8\x3c\x16\x35\x77\xa9\x13\x88\x09\xa4\xf4\xcf\x12\x99\x48\x48\xd8\x82\x10\x84\xde\x3d\xc8\x8a\x8b\x90\x6a\xa4\xd9\x6f\xd8\x59\xfa\xf9\x3f\xfe\x75\xd7\xd3\x01\x3a\x58\xa3\x04\x15\x01\x48\x73\x3f\xe5\x07\x40\x61\x0e\x86\xd5\xc0\x7c\x73\x78\xb0\x19\x9d\x65\x04\xb6\xde\xf4\x81\x49\x96\x69\x8e\xc8\x3d\x43\x83\x9b\x9e\x4a\x88\x2a\xc6\x9d\x43\x3d\x18\x37\xde\xeb\x94\x3e\xdd\x1c\xe1\xe5\xcb\xde\xe0\x55\x6e\xd2\x04\x6f\x32\x20\x2a\xd2\xad\x8a\x11\x4d\x97\x57\x22\x25\x92\xba\x55\xa2\x22\xaf\xd0\xe9\x79\xc5\xbc\xe2\xec\x2c\x5f\xfa\xe5\x17\x99\x7b\x7f\x36\xd2\x97\x8a\xa5\x19\xc9\x72\xe7\x2b\x00\x00\xff\xff\xb2\x39\xca\x9f\xff\x00\x00\x00")

func viewsIndexTplBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func viewsIndexTpl() (*asset, error) { _ = "STUB: not implemented"; return nil, nil }

func Asset(name string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MustAsset(name string) []byte { _ = "STUB: not implemented"; return nil }

func AssetInfo(name string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func AssetNames() []string { _ = "STUB: not implemented"; return nil }

var _bindata = map[string]func() (*asset, error){
	"views/blocks/block.tpl": viewsBlocksBlockTpl,
	"views/header.tpl":       viewsHeaderTpl,
	"views/index.tpl":        viewsIndexTpl,
}

func AssetDir(name string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"views": {nil, map[string]*bintree{
		"blocks": {nil, map[string]*bintree{
			"block.tpl": {viewsBlocksBlockTpl, map[string]*bintree{}},
		}},
		"header.tpl": {viewsHeaderTpl, map[string]*bintree{}},
		"index.tpl":  {viewsIndexTpl, map[string]*bintree{}},
	}},
}}

func RestoreAsset(dir, name string) error { _ = "STUB: not implemented"; return nil }

func RestoreAssets(dir, name string) error { _ = "STUB: not implemented"; return nil }

func _filePath(dir, name string) string { _ = "STUB: not implemented"; return "" }

func assetFS() *assetfs.AssetFS { _ = "STUB: not implemented"; return nil }
