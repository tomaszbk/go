// run


// https://golang.org/issue/662

package main

import "fmt"

func f() (int, int) { return 1, 2 }

func main() {
	s := fmt.Sprint(f())
	if s != "1 2" {	// with bug, was "{1 2}"
		println("BUG")
	}
}
