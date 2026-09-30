// run

// Verify that composite literals using selectors for
// embedded fields are assembled correctly.

package main

import "fmt"

type A struct {
	a int
	B
}

type B struct {
	b string
	C
}

type C struct {
	c any
}

func main() {
	eq(A{1, B{b: "foo"}}, A{a: 1, b: "foo"})
	eq(A{B: B{C: C{c: "foo"}}}, A{c: "foo"})
	eq(x, A{B: B{b: "foo"}})
}

func eq(x, y any) {
	if x != y {
		panic(fmt.Sprintf("%v != %v", x, y))
	}
}

// test global initializer
var x = A{b: "foo"}
