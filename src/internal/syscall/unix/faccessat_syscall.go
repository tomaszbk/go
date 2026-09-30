//go:build aix || linux

package unix

import "syscall"

var faccessat = syscall.Faccessat
