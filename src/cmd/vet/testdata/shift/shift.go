// This file contains tests for the suspicious shift checker.

package shift

func ShiftTest() {
	var i8 int8
	_ = i8 << 7
	_ = (i8 + 1) << 8 // ERROR ".i8 . 1. .8 bits. too small for shift of 8"
}
