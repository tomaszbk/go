// run


// Issue 18994: SSA didn't handle DOT STRUCTLIT for zero-valued
// STRUCTLIT.

package main

// large struct - not SSA-able
type T struct {
	a, b, c, d, e, f, g, h int
}

func main() {
	x := T{}.a
	if x != 0 {
		panic("FAIL")
	}
}
