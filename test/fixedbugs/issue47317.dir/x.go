// Issue 47317: ICE when calling ABI0 function via func value.

package main

func main() { F() }

func F() interface{} {
	g := G
	g(1)
	return G
}

func G(x int) [2]int
