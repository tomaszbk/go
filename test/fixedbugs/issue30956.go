// run


// Check for compile generated static data for literal
// composite struct

package main

import "fmt"

type X struct {
	V interface{}

	a int
	b int
	c int
}

func pr(x X) {
	fmt.Println(x.V)
}

func main() {
	pr(X{
		V: struct {
			A int
		}{42},
	})
}
