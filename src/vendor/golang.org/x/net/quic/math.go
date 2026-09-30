package quic

func abs[T ~int | ~int64](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
