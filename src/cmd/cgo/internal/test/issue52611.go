//go:build cgo

// Issue 52611: inconsistent compiler behaviour when compiling a C.struct.
// No runtime test; just make sure it compiles.

package cgotest

import (
	_ "cmd/cgo/internal/test/issue52611a"
	_ "cmd/cgo/internal/test/issue52611b"
)
