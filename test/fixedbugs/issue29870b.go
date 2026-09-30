// errorcheck


// Make sure we're compiling "_" functions at least enough
// to get to an error which is generated during walk.

package main

func _() {
	x := 7 // ERROR "declared and not used"
}
