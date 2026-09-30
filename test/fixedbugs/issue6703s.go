// errorcheck


// Check for cycles in a pointer method expression.

package ptrmethexpr

type T int

func (*T) pm() int {
	_ = x
	return 0
}

var x = (*T).pm // ERROR "initialization cycle|depends upon itself"
