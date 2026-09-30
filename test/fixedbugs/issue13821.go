// compile

// Issue 13821.  Compiler rejected "bool(true)" as not a constant.

package p

const (
	A = true
	B = bool(A)
	C = bool(true)
)
