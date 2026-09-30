// errorcheck -0 -m


// Make sure the go:noinline pragma makes it from a
// generic function to any of its stenciled instances.

package main

//go:noinline
func f[T any](x T) T {
	return x
}

func main() { // ERROR "can inline main"
	println(f(5))
}
