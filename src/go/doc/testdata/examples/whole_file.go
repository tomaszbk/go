package foo_test

import "fmt"

type X int

func (X) Foo() {
}

func (X) TestBlah() {
}

func (X) BenchmarkFoo() {
}

func Example() {
	fmt.Println("Hello, world!")
	// Output: Hello, world!
}
