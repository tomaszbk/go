//go:build (ppc64 || ppc64le) && linux && cgo

package cgotest

import "testing"

func TestPPC64CallStubs(t *testing.T) {
	testPPC64CallStubs(t)
}
