// errorcheck


// Issue 20185: type switching on untyped values (e.g. nil or consts)
// caused an internal compiler error.

package p

func F() {
	switch t := nil.(type) { // ERROR "cannot type switch on non-interface value|not an interface"
	default:
		_ = t
	}
}

const x = 1

func G() {
	switch t := x.(type) { // ERROR "cannot type switch on non-interface value|declared and not used|not an interface"
	default:
	}
}
