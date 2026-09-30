//go:build !unix && !windows

package pprof

import (
	"io"
)

// Stub call for platforms that don't support rusage.
func addMaxRSS(w io.Writer) {
}
