//go:build unix || js || wasip1

package main

import (
	"os"
	"syscall"
)

var signalsToIgnore = []os.Signal{os.Interrupt, syscall.SIGQUIT}
