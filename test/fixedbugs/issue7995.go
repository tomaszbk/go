// run


// Issue 7995: globals not flushed quickly enough.

package main

import "fmt"

var (
	p = 1
	q = &p
)

func main() {
	p = 50
	*q = 100
	s := fmt.Sprintln(p, *q)
	if s != "100 100\n" {
		println("BUG:", s)
	}
}
