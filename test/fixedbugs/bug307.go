// compile


// Valid program, gccgo reported an error.
// bug307.go:14:6: error: complex arguments must have identical types

package main

func main() {
	var f float64
	_ = complex(1/f, 0)
}
