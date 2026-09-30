// errorcheck


// Issue 27356: function parameter hiding built-in function results in compiler crash

package p

var a = []int{1,2,3}

func _(len int) {
	_ =  len(a) // ERROR "cannot call|expected function"
}

var cap = false
var _ = cap(a) // ERROR "cannot call|expected function"

