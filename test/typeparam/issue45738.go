// compile

package main

//go:noinline
func f[T any]() {
	x := 5
	g := func() int { return x }
	g()
}

func main() {
	f[int]()
}
