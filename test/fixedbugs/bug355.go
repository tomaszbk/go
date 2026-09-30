// run

package main

var f = func() int {
	type S int
	return 42
}

func main() {
	if f() != 42 {
		panic("BUG: bug355")
	}
}
