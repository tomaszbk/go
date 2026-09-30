// compile


// gccgo crashed compiling this file with a failed conversion to the
// alias type when constructing the composite literal.

package p

type I interface{ M() }
type A = I
type S struct {
	f A
}

func F(i I) S {
	return S{f: i}
}
