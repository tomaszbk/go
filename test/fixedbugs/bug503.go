// compile

// gccgo crashed compiling this file, due to failing to correctly emit
// the type descriptor for a named alias.

package p

type entry = struct {
	a, b, c int
}

var V entry
