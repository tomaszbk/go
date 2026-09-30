//go:build !cgo

package cgotlstest

import "testing"

func testTLS(t *testing.T) {
	t.Skip("cgo not supported")
}
