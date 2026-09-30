// compile


// Caused a gofrontend crash.

package p

func F(b []byte, i int) {
	*(*[1]byte)(b[i*2:]) = [1]byte{}
}
