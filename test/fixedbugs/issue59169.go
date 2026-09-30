// compile


// Issue 59169: caused gofrontend crash.

package p

func F(p *[]byte) {
	*(*[1]byte)(*p) = *(*[1]byte)((*p)[1:])
}
