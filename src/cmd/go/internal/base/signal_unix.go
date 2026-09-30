//go:build unix || js || wasip1

package base

import (
	"os"
	"syscall"
)

var signalsToIgnore = []os.Signal{os.Interrupt, syscall.SIGQUIT}

// SignalTrace is the signal to send to make a Go program
// crash with a stack trace.
var SignalTrace os.Signal = syscall.SIGQUIT
