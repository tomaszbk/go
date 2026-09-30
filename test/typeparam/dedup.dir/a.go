package a

//go:noinline
func F[T comparable](a, b T) bool {
	return a == b
}
