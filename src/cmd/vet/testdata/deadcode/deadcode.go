// This file contains tests for the dead code checker.

package deadcode

func _() int {
	print(1)
	return 2
	println() // ERROR "unreachable code"
	return 3
}
