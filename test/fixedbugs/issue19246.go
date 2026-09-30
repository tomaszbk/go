// run

// Issue 19246: Failed to evaluate some zero-sized values
// when converting them to interfaces.

package main

import "os"

type B struct{}

//go:noinline
func f(i interface{}) {}

func main() {
	defer func() {
		if recover() == nil {
			println("expected nil pointer dereference panic")
			os.Exit(1)
		}
	}()
	var b *B
	f(*b)
}
