// compile


// gccgo mishandled embedded methods of type aliases.

package p

type I int

func (I) M() {}

type T = struct {
	I
}

func F() {
	_ = T.M
	_ = struct { I }.M
}
