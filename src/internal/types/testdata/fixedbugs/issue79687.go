package p

func _() {
	copy(nil /* ERROR "argument must be a slice; have untyped nil" */, []byte{})
}

// test case from issue

func f() {
	var raw []byte

	copy(nil /* ERROR "argument must be a slice; have untyped nil" */, raw)
}
