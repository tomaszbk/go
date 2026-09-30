// errorcheck


// Issue 2452.

// Check that the error messages says 
//	bug378.go:17: 1 + 2 not used
// and not
//	bug378.go:17: 1 not used

package main

func main() {
	1 + 2 // ERROR "1 \+ 2 evaluated but not used|value computed is not used|is not used"
}
