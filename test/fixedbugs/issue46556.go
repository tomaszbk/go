// compile


package p

type A = interface{}
type B interface{}

// Test that embedding both anonymous and defined types is supported.
type C interface {
	A
	B
}
