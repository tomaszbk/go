// run

package main

type A struct {
	a int
	B
}

type B struct {
	b int
	C
}

type C struct {
	c int
}

func main() {
	_ = A{a: 1, b: 2}
	_ = A{a: 1, c: 3}
	_ = A{a: 1, b: 2, c: 3} // don't panic during compilation
}
