// errorcheck


// Check for cycles when calling an embedded method expression.

package embedmethexprcall

type T int

func (T) m() int {
	_ = x
	return 0
}

type E struct{ T }

var x = E.m(E{0}) // ERROR "initialization cycle|depends upon itself" 
