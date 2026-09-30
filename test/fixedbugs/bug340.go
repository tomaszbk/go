// errorcheck


// Issue 1606.

package main

func main() {
	var x interface{}
	switch t := x.(type) {
	case 0:		// ERROR "type"
		t.x = 1
		x.x = 1 // ERROR "type interface \{\}|reference to undefined field or method|interface with no methods|undefined"
	}
}
