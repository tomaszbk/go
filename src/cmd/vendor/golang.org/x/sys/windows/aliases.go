//go:build windows

package windows

import "syscall"

type Signal = syscall.Signal
type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
