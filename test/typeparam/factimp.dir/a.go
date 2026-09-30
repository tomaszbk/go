package a

func Fact[T interface{ int | int64 | float64 }](n T) T {
	if n == 1 {
		return 1
	}
	return n * Fact(n-1)
}
