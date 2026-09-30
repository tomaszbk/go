// compile


// Issue 77604: compiler crash when source and destination
// of copy are the same address.

package p

type T struct {
	a [192]byte
}

func f(x *T) {
	i := any(x)
	y := i.(*T)
	*y = *x
}
