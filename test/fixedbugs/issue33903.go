// compile


// Check that the shortcircuit pass correctly handles infinite loops.

package p

func f() {
	var p, q bool
	for {
		p = p && q
	}
}
