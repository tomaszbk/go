// errorcheck

// Check for cycles in a method call.

package methcall

type T int

func (T) m() int {
	_ = x
	return 0
}

var (
	t T
	x = t.m() // ERROR "initialization cycle|depends upon itself"
)
