// -lang=go1.22

// Check Go language version-specific errors.

//go:build go1.21

package p

func f() {
	for _ = range 10 /* ERROR "requires go1.22 or later" */ {
	}
}
