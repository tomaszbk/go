// errorcheck

// Issue 22904: Make sure the compiler emits a proper error message about
// invalid recursive types rather than crashing.

package p

type a struct{ b } // ERROR "invalid recursive type"
type b struct{ a } // GCCGO_ERROR "invalid recursive type"

var x interface{}

func f() {
	x = a{}
}
