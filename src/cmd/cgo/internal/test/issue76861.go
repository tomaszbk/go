//go:build cgo

package cgotest

// Issue 43639: No runtime test needed, make sure package
// cmd/cgo/internal/test/issue76861 compiles without error.

import _ "cmd/cgo/internal/test/issue76861"
