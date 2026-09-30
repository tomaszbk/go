// errorcheck


// Check for cycles in the method value of a pointer value returned
// from a function call.

package funcptrmethvalue

type T int

func (*T) pm() int {
	_ = x
	return 0
}

func pf() *T {
	return nil
}

var x = pf().pm // ERROR "initialization cycle|depends upon itself" 
