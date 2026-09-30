//go:build !plan9

package os

import "syscall"

type syscallErrorType = syscall.Errno

const (
	errENOSYS = syscall.ENOSYS
	errERANGE = syscall.ERANGE
	errENOMEM = syscall.ENOMEM
)
