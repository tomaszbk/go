package a

// Map calls the function f on every element of the slice s,
// returning a new slice of the results.
func Mapper[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}
