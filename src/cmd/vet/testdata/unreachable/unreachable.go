package unreachable

func _() {
	return
	println("unreachable") // ERROR "unreachable code"
}
