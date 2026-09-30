//go:build !linux && !darwin

package cryptotest

import "testing"

// BoundarySlices allocates a pair of slices of the given size.
//
// On this platform, the slices are not special.
func BoundarySlices(t *testing.T, size int) (start, end []byte) {
	return make([]byte, size), make([]byte, size)
}
