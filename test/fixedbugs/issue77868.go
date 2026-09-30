// compile

package p

type S struct {
	n int
	a [2]int
}

func f(i int) int {
	var arr [0]S

	// Accessing a zero-length array must not trigger an internal compiler error.
	// This code is invalid, but make sure that we can compile it.
	return arr[i].n
}
