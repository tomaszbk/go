//go:build cgo

// Issue 26743: typedef of uint leads to inconsistent typedefs error.
// No runtime test; just make sure it compiles.

package cgotest

import _ "cmd/cgo/internal/test/issue26743"
