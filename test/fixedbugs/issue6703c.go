// errorcheck

// Check for cycles in a method expression.

package methexpr

type T int

func (T) m() int {
	_ = x
	return 0
}

var x = T.m // ERROR "initialization cycle|depends upon itself"
