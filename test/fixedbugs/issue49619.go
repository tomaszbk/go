// build

// This testcase caused a linker crash in DWARF generation.

package main

//go:noinline
func f() any {
	var a []any
	return a[0]
}

func main() {
	f()
}
