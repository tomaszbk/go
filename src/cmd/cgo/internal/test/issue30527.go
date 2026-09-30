//go:build cgo

// Issue 30527: function call rewriting casts untyped
// constants to int because of ":=" usage.

package cgotest

import "cmd/cgo/internal/test/issue30527"

func issue30527G() {
	issue30527.G(nil)
}
