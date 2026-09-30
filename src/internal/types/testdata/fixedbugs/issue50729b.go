package p

type d[T any] struct{}
type (
	b d[a]
)

type a = func(c)
type c struct{ a }
