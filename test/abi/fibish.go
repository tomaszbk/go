// run

//go:build !wasm


package main

import "fmt"

// Test that register results are correctly returned (and passed)

//go:registerparams
//go:noinline
func f(x int) (int, int) {

	if x < 3 {
		return 0, x
	}

	a, b := f(x - 2)
	c, d := f(x - 1)
	return a + d, b + c
}

func main() {
	x := 40
	a, b := f(x)
	fmt.Printf("f(%d)=%d,%d\n", x, a, b)
}
