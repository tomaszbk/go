//go:build !plan9 && !windows

package net

import (
	"errors"
	"os"
	"syscall"
)

var (
	errOpNotSupported = syscall.EOPNOTSUPP

	abortedConnRequestErrors = []error{syscall.ECONNABORTED} // see accept in fd_unix.go
)

func isPlatformError(err error) bool {
	_, ok := err.(syscall.Errno)
	return ok
}

func samePlatformError(err, want error) bool {
	if op, ok := err.(*OpError); ok {
		err = op.Err
	}
	if sys, ok := err.(*os.SyscallError); ok {
		err = sys.Err
	}
	return err == want
}

func isENOBUFS(err error) bool {
	return errors.Is(err, syscall.ENOBUFS)
}
