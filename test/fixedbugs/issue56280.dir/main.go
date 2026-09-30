package main

import "test/a"

func main() { // ERROR "can inline main"
	a.F() // ERROR "inlining call to a.F" "inlining call to a.g\[go.shape.int\]"
}
