//go:build !linux

package net

import (
	"context"
)

func (sd *sysDialer) dialMPTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
	return sd.dialTCP(ctx, laddr, raddr)
}

func (sl *sysListener) listenMPTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error) {
	return sl.listenTCP(ctx, laddr)
}

func isUsingMultipathTCP(fd *netFD) bool {
	return false
}
