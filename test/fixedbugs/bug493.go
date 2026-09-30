// compile


// Test case that gccgo failed to compile.

package p

func F() []string {
	return []string{""}
}

var V = append(F())
