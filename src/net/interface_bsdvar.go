//go:build dragonfly || netbsd || openbsd

package net

import (
	"internal/routebsd"
	"syscall"
)

func interfaceMessages(ifindex int) ([]routebsd.Message, error) {
	return routebsd.FetchRIBMessages(syscall.NET_RT_IFLIST, ifindex)
}

// interfaceMulticastAddrTable returns addresses for a specific
// interface.
func interfaceMulticastAddrTable(ifi *Interface) ([]Addr, error) {
	// TODO(mikio): Implement this like other platforms.
	return nil, nil
}
