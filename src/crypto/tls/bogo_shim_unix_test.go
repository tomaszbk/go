//go:build unix && !wasm

package tls

import (
	"os"
	"syscall"
)

func pauseProcess() {
	pid := os.Getpid()
	process, _ := os.FindProcess(pid)
	process.Signal(syscall.SIGSTOP)
}
