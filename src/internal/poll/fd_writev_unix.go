//go:build dragonfly || freebsd || linux || netbsd

package poll

import (
	"syscall"
	"unsafe"
)

func writev(fd int, iovecs []syscall.Iovec) (uintptr, error) {
	var (
		r uintptr
		e syscall.Errno
	)
	for {
		r, _, e = syscall.Syscall(syscall.SYS_WRITEV, uintptr(fd), uintptr(unsafe.Pointer(&iovecs[0])), uintptr(len(iovecs)))
		if e != syscall.EINTR {
			break
		}
	}
	if e != 0 {
		return r, e
	}
	return r, nil
}
