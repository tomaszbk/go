//go:build !unix && !windows

package exec_test

import "os"

var (
	quitSignal os.Signal = nil
	pipeSignal os.Signal = nil
)
