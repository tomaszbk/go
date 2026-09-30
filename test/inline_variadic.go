// errorcheck -0 -m

// Test inlining of variadic functions.
// See issue #18116.

package foo

func head(xs ...string) string { // ERROR "can inline head" "leaking param: xs to result"
	return xs[0]
}

func f() string { // ERROR "can inline f"
	x := head("hello", "world") // ERROR "inlining call to head" "\.\.\. argument does not escape"
	return x
}
