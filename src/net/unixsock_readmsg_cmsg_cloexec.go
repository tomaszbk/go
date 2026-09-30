//go:build dragonfly || linux || netbsd || openbsd

package net

import "syscall"

const readMsgFlags = syscall.MSG_CMSG_CLOEXEC

func setReadMsgCloseOnExec(oob []byte) {}
