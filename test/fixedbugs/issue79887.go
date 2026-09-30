// build

package a

import "math/bits"

//go:noinline
func add(p, q, x, y uint64) uint64 {
	c := uint64(0)
	if p < q {
		c = 1
	}
	s, _ := bits.Add64(x, y, c)
	return s
}

//go:noinline
func sub(p, q, x, y uint64) uint64 {
	c := uint64(0)
	if p < q {
		c = 1
	}
	s, _ := bits.Sub64(x, y, c)
	return s
}
