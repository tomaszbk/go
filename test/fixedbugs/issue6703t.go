// errorcheck


// Check for cycles in the call of a pointer method expression.

package ptrmethexprcall

type T int

func (*T) pm() int {
	_ = x
	return 0
}

var x = (*T).pm(nil) // ERROR "initialization cycle|depends upon itself"
