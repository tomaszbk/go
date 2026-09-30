//errorcheck -0 -m -m


package p

func f() { // ERROR ""
	n, m := 100, 200
	_ = make([]byte, 1<<17)      // ERROR "too large for stack" ""
	_ = make([]byte, 100, 1<<17) // ERROR "too large for stack" ""
	_ = make([]byte, n, 1<<17)   // ERROR "too large for stack" ""

	_ = make([]byte, n)      // ERROR "does not escape"
	_ = make([]byte, 100, m) // ERROR "does not escape"
}
