//go:build !unix

package scan_test

import (
	"testing"
)

func makeMem(t testing.TB, nPages int) ([]uintptr, func()) {
	t.Skip("mmap unsupported")
	return nil, nil
}
