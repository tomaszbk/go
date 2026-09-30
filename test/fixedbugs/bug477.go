// compile


// Test multiple identical unnamed structs with methods.  This caused
// a compilation error with gccgo.

package p

type S1 struct{}

func (s S1) M() {}

type S2 struct {
	F1 struct {
		S1
	}
	F2 struct {
		S1
	}
}

type I interface {
	M()
}

func F() {
	var s2 S2
	var i1 I = s2.F1
	var i2 I = s2.F2
	_, _ = i1, i2
}
