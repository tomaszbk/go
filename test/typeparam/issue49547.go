// run


package main

import "fmt"

type foo int

func main() {
	want := "main.F[main.foo]"
	got := fmt.Sprintf("%T", F[foo]{})
	if got != want {
		fmt.Printf("want: %s, got: %s\n", want, got)
	}
}

type F[T any] struct {
}
