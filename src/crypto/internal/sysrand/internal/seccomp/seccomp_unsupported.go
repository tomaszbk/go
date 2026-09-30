//go:build !linux || !cgo

package seccomp

import "errors"

func DisableGetrandom() error {
	return errors.New("disabling getrandom is not supported on this system")
}
