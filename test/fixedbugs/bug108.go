// errorcheck


package main
func f() {
	v := 1 << 1025;		// ERROR "overflow|shift count too large"
	_ = v
}
