//go:build !plan9 && !js

package tool

import (
	"os"
	"syscall"
)

var signalsToForward = []os.Signal{syscall.SIGHUP, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM}
