//go:build !windows

package execenv

import "syscall"

// Default will return the default environment
// variables based on the process attributes
// provided.
//
// Defaults to syscall.Environ() on all platforms
// other than Windows.
func Default(sys *syscall.SysProcAttr) ([]string, error) {
	return syscall.Environ(), nil
}
