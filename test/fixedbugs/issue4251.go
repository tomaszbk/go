// errorcheck


// Issue 4251: slice with inverted range is an error.

package p

func F1(s []byte) []byte {
	return s[2:1]		// ERROR "invalid slice index|inverted slice range|invalid slice indices"
}

func F2(a [10]byte) []byte {
	return a[2:1]		// ERROR "invalid slice index|inverted slice range|invalid slice indices"
}

func F3(s string) string {
	return s[2:1]		// ERROR "invalid slice index|inverted slice range|invalid slice indices"
}
