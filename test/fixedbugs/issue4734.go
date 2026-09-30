// compile


// Caused gccgo to emit multiple definitions of the same symbol.

package p

type S1 struct{}

func (s *S1) M() {}

type S2 struct {
	F struct{ *S1 }
}

func F() {
	_ = struct{ *S1 }{}
}
