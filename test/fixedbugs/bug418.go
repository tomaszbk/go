// errorcheck

// Issue 3044.
// Multiple valued expressions in return lists.

package p

func Two() (a, b int)

// F used to compile.
func F() (x interface{}, y int) {
	return Two(), 0 // ERROR "single-value context|2\-valued"
}

// Recursive used to trigger an internal compiler error.
func Recursive() (x interface{}, y int) {
	return Recursive(), 0 // ERROR "single-value context|2\-valued"
}
