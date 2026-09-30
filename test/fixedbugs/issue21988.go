// errorcheck

// Issue 21988: panic on switch case with invalid value

package p

const X = Wrong(0) // ERROR "undefined: Wrong|undefined name .*Wrong"

func _() {
	switch 0 {
	case X:
	}
}
