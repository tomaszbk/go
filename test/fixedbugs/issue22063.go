// errorcheck

// Issue 22063: panic on interface switch case with invalid name

package p

const X = Wrong(0) // ERROR "undefined: Wrong|reference to undefined name .*Wrong"

func _() {
	switch interface{}(nil) {
	case X:
	}
}
