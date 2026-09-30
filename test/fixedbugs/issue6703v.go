// errorcheck


// Check for cycles in a pointer literal's method call.

package ptrlitmethcall

type T int

func (*T) pm() int {
	_ = x
	return 0
}

var x = (*T)(nil).pm() // ERROR "initialization cycle|depends upon itself"
