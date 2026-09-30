//go:build unix && !wasip1

package unix

import (
	"syscall"
	_ "unsafe" // for //go:linkname
)

//go:linkname Utimensat syscall.utimensat
func Utimensat(dirfd int, path string, times *[2]syscall.Timespec, flag int) error
