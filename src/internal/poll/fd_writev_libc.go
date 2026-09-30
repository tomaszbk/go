//go:build aix || darwin || openbsd || solaris

package poll

import (
	"syscall"
	_ "unsafe" // for go:linkname
)

//go:linkname writev syscall.writev
func writev(fd int, iovecs []syscall.Iovec) (uintptr, error)
