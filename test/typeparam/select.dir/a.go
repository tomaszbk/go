package a

func F[T any](c, d chan T) T {
	select {
	case x := <- c:
		return x
	case x := <- d:
		return x
	}
}

