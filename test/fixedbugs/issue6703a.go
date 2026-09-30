// errorcheck


// Check for cycles in a function value.

package funcvalue

func fx() int {
	_ = x
	return 0
}

var x = fx // ERROR "initialization cycle|depends upon itself"
