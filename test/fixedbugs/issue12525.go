// errorcheck

// Issue 12525: confusing error trying to increment boolean value

package main

func main() {
	var i int
	i++

	var f float64
	f++

	var c complex128
	c++

	var b bool
	b++ // ERROR "invalid operation: b\+\+ \(non-numeric type bool\)"

	var s string
	s-- // ERROR "invalid operation: s-- \(non-numeric type string\)"
}
