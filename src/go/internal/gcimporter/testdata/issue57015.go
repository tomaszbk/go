package issue57015

type E error

type X[T any] struct {}

func F() X[interface {
	E
}] {
	panic(0)
}

