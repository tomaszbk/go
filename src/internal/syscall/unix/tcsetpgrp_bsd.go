//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package unix

import (
	"syscall"
	"unsafe"
)

//go:linkname ioctlPtr syscall.ioctlPtr
func ioctlPtr(fd int, req uint, arg unsafe.Pointer) (err error)

// Note that pgid should really be pid_t, however _C_int (aka int32) is
// generally equivalent.

func Tcsetpgrp(fd int, pgid int32) (err error) {
	return ioctlPtr(fd, syscall.TIOCSPGRP, unsafe.Pointer(&pgid))
}
