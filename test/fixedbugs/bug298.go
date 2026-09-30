// errorcheck


package ddd

func Sum() int
	for i := range []int{} { return i }  // ERROR "statement outside function|expected"

