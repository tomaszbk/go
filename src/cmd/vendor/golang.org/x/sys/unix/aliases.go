//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

package unix

import "syscall"

type Signal = syscall.Signal
type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
