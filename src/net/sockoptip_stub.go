//go:build js || wasip1

package net

import "syscall"

func setIPv4MulticastInterface(fd *netFD, ifi *Interface) error {
	return syscall.ENOPROTOOPT
}

func setIPv4MulticastLoopback(fd *netFD, v bool) error {
	return syscall.ENOPROTOOPT
}

func joinIPv4Group(fd *netFD, ifi *Interface, ip IP) error {
	return syscall.ENOPROTOOPT
}

func setIPv6MulticastInterface(fd *netFD, ifi *Interface) error {
	return syscall.ENOPROTOOPT
}

func setIPv6MulticastLoopback(fd *netFD, v bool) error {
	return syscall.ENOPROTOOPT
}

func joinIPv6Group(fd *netFD, ifi *Interface, ip IP) error {
	return syscall.ENOPROTOOPT
}
