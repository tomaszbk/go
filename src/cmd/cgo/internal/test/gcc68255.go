//go:build cgo

package cgotest

import (
	"testing"

	"cmd/cgo/internal/test/gcc68255"
)

func testGCC68255(t *testing.T) {
	if !gcc68255.F() {
		t.Error("C global variable was not initialized")
	}
}
