package foo_test

// Variable declaration with fewer values than names.

func f() (int, int) {
	return 1, 2
}

var a, b = f()

// Need two examples to hit playExample.

func ExampleA() {
	_ = a
}

func ExampleB() {
}
