// run


// Verify that generic methods work with pointer receivers.

package main

import "fmt"

type T struct {}

func (t *T) m[P any]() {
	var x P
	if got := fmt.Sprintf("%T", x); got != "int" {
		panic(fmt.Sprintf("got %s, want int", got))
	}
}

type G[P any] struct {}

func (g *G[P]) m[Q any]() {
	var p P
	var q Q
	if got := fmt.Sprintf("%T %T", p, q); got != "int bool" {
		panic(fmt.Sprintf("got %s, want int bool", got))
	}
}

func main() {
	(&T{}).m[int]()
	(&G[int]{}).m[bool]()
}
