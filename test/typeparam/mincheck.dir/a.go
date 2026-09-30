package a

type Ordered interface {
	int | int64 | float64
}

func Min[T Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
