//go:build unix || js || wasip1

package net

import "syscall"

func isConnError(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.ECONNRESET || se == syscall.ECONNABORTED
	}
	return false
}
