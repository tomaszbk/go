// run

package main

import (
	"fmt"
)

type foo struct {
	bar [1]*int
}

func main() {
	ch := make(chan foo, 2)
	var a int
	var b [1]*int
	b[0] = &a
	ch <- foo{bar: b}
	close(ch)

	for v := range ch {
		for i := 0; i < 1; i++ {
			fmt.Println(v.bar[0] != nil)
		}
	}
}
