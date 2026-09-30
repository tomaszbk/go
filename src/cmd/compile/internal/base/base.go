package base

import (
	"os"
)

var atExitFuncs []func()

func AtExit(f func()) {
	atExitFuncs = append(atExitFuncs, f)
}

func Exit(code int) {
	for i := len(atExitFuncs) - 1; i >= 0; i-- {
		f := atExitFuncs[i]
		atExitFuncs = atExitFuncs[:i]
		f()
	}
	os.Exit(code)
}

// To enable tracing support (-t flag), set EnableTrace to true.
const EnableTrace = false
