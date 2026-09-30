package b

import "./a"

var x a.Foo

func main() {
	x.int = 20    // ERROR "unexported field|undefined"
	x.int8 = 20   // ERROR "unexported field|undefined"
	x.error = nil // ERROR "unexported field|undefined"
	x.rune = 'a'  // ERROR "unexported field|undefined"
	x.byte = 20   // ERROR "unexported field|undefined"
}
