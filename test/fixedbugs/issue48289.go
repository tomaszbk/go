// run

package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	var ptrs [2]*int
	for i := range ptrs {
		ch <- i
		select {
		case x := <-ch:
			ptrs[i] = &x
		}
	}

	for i, ptr := range ptrs {
		if *ptr != i {
			panic(fmt.Sprintf("got *ptr %d, want %d", *ptr, i))
		}
	}
}
