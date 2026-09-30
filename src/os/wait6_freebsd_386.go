package os

import (
	"syscall"
	"unsafe"
)

const _P_PID = 0

func wait6(idtype, id, options int) (status int, errno syscall.Errno) {
	// freebsd32_wait6_args{ idtype, id1, id2, status, options, wrusage, info }
	_, _, errno = syscall.Syscall9(syscall.SYS_WAIT6, uintptr(idtype), uintptr(id), 0, uintptr(unsafe.Pointer(&status)), uintptr(options), 0, 0, 0, 0)
	return status, errno
}
