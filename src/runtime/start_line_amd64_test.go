package runtime_test

import (
	"internal/runtime/startlinetest"
	"testing"
)

// TestStartLineAsm tests the start line metadata of an assembly function. This
// is only tested on amd64 to avoid the need for a proliferation of per-arch
// copies of this function.
func TestStartLineAsm(t *testing.T) {
	startlinetest.CallerStartLine = callerStartLine

	const wantLine = 23
	got := startlinetest.AsmFunc()
	if got != wantLine {
		t.Errorf("start line got %d want %d", got, wantLine)
	}
}
