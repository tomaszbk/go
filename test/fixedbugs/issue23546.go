// run


// Issue 23546: type..eq function not generated when
// DWARF is disabled.

package main

func main() {
	use(f() == f())
}

func f() [2]interface{} {
	var out [2]interface{}
	return out
}

//go:noinline
func use(bool) {}
