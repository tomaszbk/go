//go:build !(cgo && (aix || dragonfly || freebsd || (linux && !android) || netbsd || openbsd)) && !darwin

package testpty

import "os"

func open() (pty *os.File, processTTY string, err error) {
	return nil, "", ErrNotSupported
}
