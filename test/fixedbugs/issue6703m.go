// errorcheck

// Check for cycles in the method value of a value returned from a function call.

package funcmethvalue

type T int

func (T) m() int {
	_ = x
	return 0
}

func f() T {
	return T(0)
}

var (
	t T
	x = f().m // ERROR "initialization cycle|depends upon itself"
)
