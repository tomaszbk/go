package p

func _() {
M:
L:
	for range 0 {
		break L
		break /* ERROR invalid break label M */ M
	}
	for range 0 {
		break /* ERROR invalid break label L */ L
	}
}
