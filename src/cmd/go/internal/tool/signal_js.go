//go:build js

package tool

import (
	"os"
	"syscall"
)

var signalsToForward = []os.Signal{os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM}
