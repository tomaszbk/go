// run

// Test that generic interface-interface comparisons resulting from
// value switch statements are handled correctly.

package main

func main() {
	f[X](0)
}

type Mer[T any] interface{ M(T) }
type MNer[T any] interface {
	Mer[T]
	N()
}

type X int

func (X) M(X) {}
func (X) N()  {}

func f[T MNer[T]](t T) {
	switch Mer[T](t) {
	case MNer[T](t):
		// ok
	default:
		panic("FAIL")
	}
}
