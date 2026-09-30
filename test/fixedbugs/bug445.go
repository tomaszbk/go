// compile

// Issue 3765

package main

func f(x uint) uint {
	m := ^(1 << x)
	return uint(m)
}
