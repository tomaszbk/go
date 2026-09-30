// errorcheck

// Don't crash while reporting the error.

package p

func _() {
	type T = T // ERROR "invalid recursive type: T refers to itself"
}
