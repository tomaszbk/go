package net

import (
	"internal/syscall/windows"
	"syscall"
)

var (
	hostsFilePath = windows.GetSystemDirectory() + "/Drivers/etc/hosts"

	// Placeholders for socket system calls.
	wsaSocketFunc func(int32, int32, int32, *syscall.WSAProtocolInfo, uint32, uint32) (syscall.Handle, error) = windows.WSASocket
	connectFunc   func(syscall.Handle, syscall.Sockaddr) error                                                = syscall.Connect
	listenFunc    func(syscall.Handle, int) error                                                             = syscall.Listen
)
