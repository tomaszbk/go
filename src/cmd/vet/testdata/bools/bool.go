// This file contains tests for the bool checker.

package bool

func _() {
	var f, g func() int

	if v, w := f(), g(); v == w || v == w { // ERROR "redundant or: v == w || v == w"
	}
}
