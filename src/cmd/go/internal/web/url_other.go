//go:build !windows

package web

import (
	"errors"
	"path/filepath"
)

func convertFileURLPath(host, path string) (string, error) {
	switch host {
	case "", "localhost":
	default:
		return "", errors.New("file URL specifies non-local host")
	}
	return filepath.FromSlash(path), nil
}
