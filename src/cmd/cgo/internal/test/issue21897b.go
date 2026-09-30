//go:build !darwin || !cgo || internal

package cgotest

import "testing"

func test21897(t *testing.T) {
	t.Skip("test runs only on darwin+cgo")
}
