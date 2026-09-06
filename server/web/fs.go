package web

import (
	"net/http"
	"os"
	"path/filepath"
)

type FileSystem struct{}

func (d FileSystem) Open(name string) (http.File, error) {
	_ = "STUB: not implemented"
	return *new(http.File), nil
}

func Walk(fs http.FileSystem, root string, walkFn filepath.WalkFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func walk(fs http.FileSystem, path string, info os.FileInfo, walkFn filepath.WalkFunc) error {
	_ = "STUB: not implemented"
	return nil
}
