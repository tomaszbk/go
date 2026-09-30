// errorcheck

package main

type bar struct {
	x int
}

func main() {
	var foo bar
	_ = &foo{} // ERROR "is not a type|expected .;."
} // GCCGO_ERROR "expected declaration"
