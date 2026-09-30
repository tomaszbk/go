//go:build unix

package exec_test

import (
	"os"
	"syscall"
)

var (
	quitSignal os.Signal = syscall.SIGQUIT
	pipeSignal os.Signal = syscall.SIGPIPE
)
