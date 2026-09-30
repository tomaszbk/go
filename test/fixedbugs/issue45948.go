// compile -N


// Issue 45948: assert in debug generation for degenerate
// function with infinite loop.

package p

func f(p int) {
L:
	goto L

}
