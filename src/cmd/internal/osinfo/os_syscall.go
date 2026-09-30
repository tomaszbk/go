//go:build aix || linux

// Supporting definitions for os_uname.go on AIX and Linux.

package osinfo

import "syscall"

type utsname = syscall.Utsname

func uname(buf *utsname) error {
	return syscall.Uname(buf)
}
