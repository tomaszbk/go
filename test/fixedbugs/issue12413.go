// compile


// issue 12413: invalid variable name x in type switch: code would fail
// to compile if the variable used in the short variable declaration was
// previously declared as a constant.

package main

func main() {
	const x = 42
	switch x := interface{}(nil).(type) {
	default:
		_ = x
	}
}
