// run

//go:build !wasm

package main

//go:noinline
func F() {
	b := g()
	defer g2(b)
	n := g()[20]
	println(n)
}

type T [45]int

var x = 0

//go:noinline
func g() T {
	x++
	return T{20: x}
}

//go:noinline
func g2(t T) {
	if t[20] != 1 {
		println("FAIL", t[20])
	}
}

func main() { F() }
