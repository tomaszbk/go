// run


// Test that implicit conversions to interface type in a select/case
// clause are compiled correctly.

package main

import "fmt"

func main() { f[int]() }

func f[T any]() {
	ch := make(chan T)
	close(ch)

	var i, ok any
	select {
	case i, ok = <-ch:
	}

	fmt.Printf("%T %T\n", i, ok)
}
