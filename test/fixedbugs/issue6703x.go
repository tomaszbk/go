// errorcheck

// Check for cycles in a pointer value's method call.

package ptrmethcall

type T int

func (*T) pm() int {
	_ = x
	return 0
}

var (
	p *T
	x = p.pm() // ERROR "initialization cycle|depends upon itself"
)
