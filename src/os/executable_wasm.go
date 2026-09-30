//go:build wasm

package os

import (
	"errors"
	"runtime"
)

func executable() (string, error) {
	return "", errors.New("Executable not implemented for " + runtime.GOOS)
}
