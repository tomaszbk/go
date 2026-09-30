//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

package poll

import "syscall"

// Accept4Func is used to hook the accept4 call.
var Accept4Func func(int, int) (int, syscall.Sockaddr, error) = syscall.Accept4
