// errorcheck


// Issue 7538: blank (_) labels handled incorrectly

package p

func f() {
_:
_:
	goto _ // ERROR "not defined|undefined label"
}
