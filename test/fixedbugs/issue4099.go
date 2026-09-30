// errorcheck -0 -m

//go:build !goexperiment.newinliner

// Check go:noescape annotations.

package p

// The noescape comment only applies to the next func,
// which must not have a body.

//go:noescape

func F1([]byte)

func F2([]byte)

func G() {
	var buf1 [10]byte
	F1(buf1[:])

	var buf2 [10]byte // ERROR "moved to heap: buf2"
	F2(buf2[:])
}
