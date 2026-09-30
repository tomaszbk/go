// errorcheck

// Verify that pointer method calls are caught during typechecking.
// Reproducer extracted and adapted from method.go

package foo

type A struct {
	B
}
type B int

func (*B) g() {}

var _ = func() {
	var a A
	A(a).g() // ERROR "cannot call pointer method .*on|cannot take the address of"
}
