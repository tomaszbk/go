// errorcheck


// Issue 4470: parens are not allowed around .(type) "expressions"

package main

func main() {
	var i interface{}
	switch (i.(type)) { // ERROR "outside type switch"
	default:
	}
	_ = i
}
