// errorcheck


package main

var x int

func main() {
	(x) := 0  // ERROR "non-name [(]x[)]|non-name on left side"
}
