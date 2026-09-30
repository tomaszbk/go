// compile


// Issue 17551: inrange optimization failed to preserve type information.

package main

import "fmt"

func main() {
	_, x := X()
	fmt.Printf("x = %v\n", x)
}

func X() (i int, ok bool) {
	ii := int(1)
	return ii, 0 <= ii && ii <= 0x7fffffff
}
