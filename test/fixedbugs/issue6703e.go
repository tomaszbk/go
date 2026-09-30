// errorcheck


// Check for cycles in the method value of a value literal.

package litmethvalue

type T int

func (T) m() int {
	_ = x
	return 0
}

var x = T(0).m // ERROR "initialization cycle|depends upon itself"
