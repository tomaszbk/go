// compile -d=ssa/check/seed

// This code caused an internal consistency error due to a bad shortcircuit optimization.

package p

func f() {
	var b bool
	if b {
		b = true
	}
l:
	for !b {
		b = true
		goto l
	}
}
