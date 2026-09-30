//go:build js

package net

import (
	"os"
	"syscall"
)

func fileConn(f *os.File) (Conn, error)             { return nil, syscall.ENOPROTOOPT }
func fileListener(f *os.File) (Listener, error)     { return nil, syscall.ENOPROTOOPT }
func filePacketConn(f *os.File) (PacketConn, error) { return nil, syscall.ENOPROTOOPT }
