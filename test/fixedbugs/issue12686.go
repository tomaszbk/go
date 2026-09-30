// compile


// golang.org/issue/12686.
// interesting because it's a non-constant but ideal value
// and we used to incorrectly attach a constant Val to the Node.

package p

func f(i uint) uint {
	x := []uint{1 << i}
	return x[0]
}
