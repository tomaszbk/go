//go:build !linux

package poll_test

import (
	"errors"
	"os"
	"runtime"
)

func badStateFile() (*os.File, error) {
	return nil, errors.New("not supported on " + runtime.GOOS)
}

func isBadStateFileError(err error) (string, bool) {
	return "", false
}
