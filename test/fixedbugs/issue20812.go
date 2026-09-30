// errorcheck

package p

func f() {
	_ = int("1")      // ERROR "cannot convert|invalid type conversion"
	_ = bool(0)       // ERROR "cannot convert|invalid type conversion"
	_ = bool("false") // ERROR "cannot convert|invalid type conversion"
	_ = int(false)    // ERROR "cannot convert|invalid type conversion"
	_ = string(true)  // ERROR "cannot convert|invalid type conversion"
}
