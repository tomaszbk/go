// compile

// Issue 1811.
// gccgo failed to compile this.

package p

type E interface{}

type I interface {
	E
	E
}
