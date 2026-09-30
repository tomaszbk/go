//go:build !plan9 && !windows

package exec

import (
	"io/fs"
	"syscall"
)

// skipStdinCopyError optionally specifies a function which reports
// whether the provided stdin copy error should be ignored.
func skipStdinCopyError(err error) bool {
	// Ignore EPIPE errors copying to stdin if the program
	// completed successfully otherwise.
	// See Issue 9173.
	pe, ok := err.(*fs.PathError)
	return ok &&
		pe.Op == "write" && pe.Path == "|1" &&
		pe.Err == syscall.EPIPE
}
