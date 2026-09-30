// errorcheck


// issue 8311.
// error for x++ should say x++ not x += 1

package p

func f() {
	var x []byte
	x++ // ERROR "invalid operation: x[+][+]|non-numeric type"

}
