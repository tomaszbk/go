// run


package main

// Make sure the compiler knows that DUFFCOPY clobbers X0

import "fmt"

//go:noinline
func f(x float64) float64 {
	// y is allocated to X0
	y := x + 5
	// marshals z before y.  Marshaling z
	// calls DUFFCOPY.
	return g(z, y)
}

//go:noinline
func g(b [64]byte, y float64) float64 {
	return y
}

var z [64]byte

func main() {
	got := f(5)
	if got != 10 {
		panic(fmt.Sprintf("want 10, got %f", got))
	}
}
