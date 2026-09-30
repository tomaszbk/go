// errorcheck

// Check for cycles in an embedded struct literal's method call.

package embedlitmethcall

type T int

func (T) m() int {
	_ = x
	return 0
}

type E struct{ T }

var x = E{}.m() // ERROR "initialization cycle|depends upon itself"
