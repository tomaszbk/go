// -lang=go1.21

// Check Go language version-specific errors.

//go:build go1.22

package p

func f() {
	for _ = range /* ok because of upgrade to 1.22 */ 10 {
	}
}
