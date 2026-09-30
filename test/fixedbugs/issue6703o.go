// errorcheck


// Check for cycles in an embedded struct's method value.

package embedmethvalue

type T int

func (T) m() int {
	_ = x
	return 0
}

type E struct{ T }

var (
	e E
	x = e.m // ERROR "initialization cycle|depends upon itself" 
)
