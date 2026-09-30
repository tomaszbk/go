// compile


// Gccgo used to incorrectly give an error when compiling this.

package p

func F() (i int) {
	for first := true; first; first = false {
		i++
	}
	return
}
