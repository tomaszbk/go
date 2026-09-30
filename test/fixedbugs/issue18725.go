// run


package main

import "os"

func panicWhenNot(cond bool) {
	if cond {
		os.Exit(0)
	} else {
		panic("nilcheck elim failed")
	}
}

func main() {
	e := (*string)(nil)
	panicWhenNot(e == e)
	// Should never reach this line.
	panicWhenNot(*e == *e)
}
