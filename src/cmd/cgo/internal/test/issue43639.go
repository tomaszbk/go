//go:build cgo

package cgotest

// Issue 43639: No runtime test needed, make sure package cmd/cgo/internal/test/issue43639 compiles well.

import _ "cmd/cgo/internal/test/issue43639"
