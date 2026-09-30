// errorcheck


//go:build amd64 || arm64

// Issue 78355: map element or key type too large should not cause ICE.

package p

type T [1 << 31]byte

func F(m map[int]T) { // ERROR "map element type too large"
	_ = m[0]
}
