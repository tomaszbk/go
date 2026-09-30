//go:build darwin

package net

import (
	"syscall"
	"testing"
)

const (
	syscall_TCP_KEEPIDLE  = syscall.TCP_KEEPALIVE
	syscall_TCP_KEEPCNT   = sysTCP_KEEPCNT
	syscall_TCP_KEEPINTVL = sysTCP_KEEPINTVL
)

type fdType = int

func maybeSkipKeepAliveTest(_ *testing.T) {}
