//go:build !windows

package cgotest

import "testing"

func test8517(t *testing.T) {
	t.Skip("skipping windows only test")
}
