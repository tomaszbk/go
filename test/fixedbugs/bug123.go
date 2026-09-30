// errorcheck


package main
const ( F = 1 )
func fn(i int) int {
	if i == F() {		// ERROR "func"
		return 0
	}
	return 1
}
