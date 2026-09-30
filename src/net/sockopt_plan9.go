package net

import "syscall"

func setKeepAlive(fd *netFD, keepalive bool) error {
	if keepalive {
		_, e := fd.ctl.WriteAt([]byte("keepalive"), 0)
		return e
	}
	return nil
}

func setLinger(fd *netFD, sec int) error {
	return syscall.EPLAN9
}
