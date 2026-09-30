// errorcheck


package p

func f() bool {
	if false {
	} else {
		return true
	}
} // ERROR "missing return( at end of function)?"
