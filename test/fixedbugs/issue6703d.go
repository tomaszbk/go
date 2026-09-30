// errorcheck


// Check for cycles in a method expression call.

package methexprcall

type T int

func (T) m() int {
	_ = x
	return 0
}

var x = T.m(0) // ERROR "initialization cycle|depends upon itself"
