//go:build unix

package cgotest

import (
	"syscall"
	"testing"
)

var syscall_dot_SIGCHLD = syscall.SIGCHLD

func usesUCRT(t *testing.T) bool {
	return false
}
