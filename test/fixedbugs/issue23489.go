// run

// Caused gccgo to issue a spurious compilation error.

package main

type T struct{}

func (*T) Foo() {}

type P = *T

func main() {
	var p P
	p.Foo()
}
