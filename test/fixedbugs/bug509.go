// compile

// Gccgo mishandles a couple of alias cases.

package p

type S struct{}

func (*S) M() {}

type I interface {
	M()
}

type A = *S

var V1 I
var _ = V1.(*S)
var _ = V1.(A)

func F() {
	var v I
	v = (*S)(nil)
	v = A(nil)
	_ = v
}
