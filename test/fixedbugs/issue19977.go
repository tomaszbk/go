// errorcheck

// Issue 19977: multiple error messages when type switching on an undefined

package foo

func Foo() {
	switch x := a.(type) { // ERROR "undefined: a|reference to undefined name .*a"
	default:
		_ = x
	}
}
