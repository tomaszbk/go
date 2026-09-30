//line range_esc_closure_linedir.go:5
package main

import "fmt"

var is []func() int

func main() {
	var ints = []int{0, 0, 0}
	for i := range ints {
		is = append(is, func() int { return i })
	}

	for _, f := range is {
		fmt.Println(f())
		if f() != 2 {
			panic("loop variable i: expected shared per-loop, but got distinct per-iteration")
		}
	}
}
