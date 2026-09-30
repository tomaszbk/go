package unix

import (
	"syscall"
	"unsafe"
)

// Note that pgid should really be pid_t, however _C_int (aka int32) is
// generally equivalent.

func Tcsetpgrp(fd int, pgid int32) (err error) {
	_, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCSPGRP), uintptr(unsafe.Pointer(&pgid)), 0, 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}
