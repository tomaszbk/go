package p

type myString string

func _[P ~string | ~[]byte | ~[]rune]() {
	_ = P("")
	const s myString = ""
	_ = P(s)
}

func _[P myString]() {
	_ = P("")
}
