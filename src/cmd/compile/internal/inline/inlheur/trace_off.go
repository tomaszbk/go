//go:build !debugtrace

package inlheur

const debugTrace = 0

func enableDebugTrace(x int) {
}

func enableDebugTraceIfEnv() {
}

func disableDebugTrace() {
}
