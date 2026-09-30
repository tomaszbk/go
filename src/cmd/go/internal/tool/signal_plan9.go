//go:build plan9

package tool

import (
	"os"
	"syscall"
)

var signalsToForward = []os.Signal{syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}
