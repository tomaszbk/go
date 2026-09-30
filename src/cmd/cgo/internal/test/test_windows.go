package cgotest

import (
	"internal/syscall/windows"
	"syscall"
	"testing"
)

var syscall_dot_SIGCHLD syscall.Signal

// usesUCRT reports whether the test is using the Windows UCRT (Universal C Runtime).
func usesUCRT(t *testing.T) bool {
	name, err := syscall.UTF16PtrFromString("ucrtbase.dll")
	if err != nil {
		t.Fatal(err)
	}
	h, err := windows.GetModuleHandle(name)
	return err == nil && h != 0
}
