// errorcheck


// Verify that error message reports _ambiguous_ method.

package p

type A struct{
	H int
}

func (A) F() {}
func (A) G() {}

type B struct{
	G int
	H int
}

func (B) F() {}

type C struct {
	A
	B
}

var _ = C.F // ERROR "ambiguous"
var _ = C.G // ERROR "ambiguous"
var _ = C.H // ERROR "ambiguous"
var _ = C.I // ERROR "no method .*I.*|C.I undefined"
