//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package syscall

import (
	"unsafe"
)

// pgid should really be pid_t, however _C_int (aka int32) is generally
// equivalent.

func Tcgetpgrp(fd int) (pgid int32, err error) {
	if err := ioctlPtr(fd, TIOCGPGRP, unsafe.Pointer(&pgid)); err != nil {
		return -1, err
	}
	return pgid, nil
}

func Tcsetpgrp(fd int, pgid int32) (err error) {
	return ioctlPtr(fd, TIOCSPGRP, unsafe.Pointer(&pgid))
}
