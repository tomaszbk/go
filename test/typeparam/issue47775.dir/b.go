package b

type C[T any] struct {
}

func (c *C[T]) reset() {
}

func New[T any]() {
	c := &C[T]{}
	z(c.reset)
}

func z(interface{}) {
}
