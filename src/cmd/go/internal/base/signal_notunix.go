//go:build plan9 || windows

package base

import (
	"os"
)

var signalsToIgnore = []os.Signal{os.Interrupt}

// SignalTrace is the signal to send to make a Go program
// crash with a stack trace (no such signal in this case).
var SignalTrace os.Signal = nil
