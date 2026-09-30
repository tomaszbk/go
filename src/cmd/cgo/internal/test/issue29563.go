//go:build cgo && !windows

// Issue 29563: internal linker fails on duplicate weak symbols.
// No runtime test; just make sure it compiles.

package cgotest

import _ "cmd/cgo/internal/test/issue29563"
