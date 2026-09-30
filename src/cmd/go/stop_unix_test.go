//go:build unix || (js && wasm)

package main_test

import (
	"os"
	"syscall"
)

func quitSignal() os.Signal {
	return syscall.SIGQUIT
}
