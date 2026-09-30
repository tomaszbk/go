package p

func _[T bool](ch chan T) {
	var _, _ T = <-ch
}

// offending code snippets from issue

func _[T ~bool](ch <-chan T) {
	var x, ok T = <-ch
	println(x, ok)
}

func _[T ~bool](m map[int]T) {
	var x, ok T = m[0]
	println(x, ok)
}
