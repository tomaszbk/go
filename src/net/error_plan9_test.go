package net

import "syscall"

var (
	errOpNotSupported = syscall.EPLAN9

	abortedConnRequestErrors []error
)

func isPlatformError(err error) bool {
	_, ok := err.(syscall.ErrorString)
	return ok
}

func isENOBUFS(err error) bool {
	return false // ENOBUFS is Unix-specific
}
