//go:build !windows

package cgotest

import "testing"

func test42018(t *testing.T) {
	t.Skip("skipping Windows-only test")
}
