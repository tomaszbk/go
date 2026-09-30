// -lang=go1.25

// Check Go language version-specific errors.

//go:build go1.25

package p

func f(x int) {
	_ = new /* ERROR "new(123) requires go1.26 or later" */ (123)
	_ = new /* ERROR "new(x) requires go1.26 or later" */ (x)
	_ = new /* ERROR "new(f) requires go1.26 or later" */ (f)
	_ = new /* ERROR "new(1 < 2) requires go1.26 or later" */ (1 < 2)
}
