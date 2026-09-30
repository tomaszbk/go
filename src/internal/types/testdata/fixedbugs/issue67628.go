package p

func f[P any](x P) P { return x }

func _() {
	type A = int
	var a A
	b := f(a) // type of b is A
	// error should report type of b as A, not int
	_ = b /* ERROR "mismatched types A and untyped string" */ + "foo"
}
