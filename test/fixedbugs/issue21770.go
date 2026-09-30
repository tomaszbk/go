// errorcheck

// Issue 21770: gccgo incorrectly accepts "p.f = 0" where p is **struct

package p

type PP **struct{ f int }

func f() {
	// anonymous type
	var p **struct{ f int }
	p.f = 0 // ERROR "field"
	// named type
	var p2 PP
	p2.f = 0 // ERROR "field"
}
