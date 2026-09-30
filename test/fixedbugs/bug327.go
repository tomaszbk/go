// run


// Conversion between identical interfaces.
// Issue 1647.

// The compiler used to not realize this was a no-op,
// so it generated a call to the non-existent function runtime.convE2E.

package main

type (
	a interface{}
	b interface{}
)

func main() {
	x := a(1)
	z := b(x)
	_ = z
}
