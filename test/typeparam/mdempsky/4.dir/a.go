package a

func F[T any](T) {
Loop:
	for {
		break Loop
	}
}
