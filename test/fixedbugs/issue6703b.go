// errorcheck

// Check for cycles in a function call.

package funccall

func fx() int {
	_ = x
	return 0
}

var x = fx() // ERROR "initialization cycle|depends upon itself"
