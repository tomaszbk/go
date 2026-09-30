// compile

// Issue 6247: 5g used to be confused by the numbering
// of floating-point registers.

package main

var p map[string]interface{}
var v interface{}

func F() {
	p["hello"] = v.(complex128) * v.(complex128)
}
