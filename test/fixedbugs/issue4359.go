// errorcheck

// Issue 4359: wrong handling of broken struct fields
// causes "internal compiler error: lookdot badwidth".

package main

type T struct {
	x T1 // ERROR "undefined"
}

func f() {
	var t *T
	_ = t.x
}
