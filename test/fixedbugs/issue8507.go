// errorcheck

// issue 8507
// used to call algtype on invalid recursive type and get into infinite recursion

package p

type T struct{ T } // ERROR "invalid recursive type.*T"

func f() {
	println(T{} == T{})
}
