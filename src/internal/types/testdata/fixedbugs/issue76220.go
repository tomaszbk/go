package p

func _() {
	append(nil /* ERROR "argument must be a slice; have untyped nil" */, ""...)
}

// test case from issue

func main() {
	s := "hello"
	msg := append(nil /* ERROR "argument must be a slice; have untyped nil" */, s...)
	print(msg)
}
