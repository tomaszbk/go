// compile

// Logical operation on named boolean type returns the same type,
// supporting an implicit conversion to an interface type.  This used
// to crash gccgo.

package p

type B bool

func (b B) M() {}

type I interface {
	M()
}

func F(a, b B) I {
	return a && b
}
