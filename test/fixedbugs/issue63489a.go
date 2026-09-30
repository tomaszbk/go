// errorcheck -lang=go1.22

// This file has been changed from its original version as
// //go:build file versions below 1.21 set the language version to 1.21.
// The original tested a -lang version of 1.21 with a file version of
// go1.4 while this new version tests a -lang version of go1.22
// with a file version of go1.21.

//go:build go1.21

package p

func f() {
	for _ = range 10 { // ERROR "file declares //go:build go1.21"
	}
}
