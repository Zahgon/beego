//go:build !windows
// +build !windows

package utils

import (
	"os"
)

func OpenFileSecure(name string, flag int, perm os.FileMode) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
