//go:build unix || js || wasip1

package net

import "syscall"

var (
	testHookCanceledDial = func() {} // for golang.org/issue/16523

	hostsFilePath = "/etc/hosts"

	// Placeholders for socket system calls.
	socketFunc        func(int, int, int) (int, error)  = syscall.Socket
	connectFunc       func(int, syscall.Sockaddr) error = syscall.Connect
	listenFunc        func(int, int) error              = syscall.Listen
	getsockoptIntFunc func(int, int, int) (int, error)  = syscall.GetsockoptInt
)
