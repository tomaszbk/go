// run


package main

import "fmt"

var indent uint = 10
func main() {
	const dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . " +
		". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
	const n = uint(len(dots))
	i := 2 * indent
	var s string
	for ; i > n; i -= n {
		s += fmt.Sprint(dots)
	}
	s += dots[0:i]
	if s != ". . . . . . . . . . " {
		panic(s)
	}
}
