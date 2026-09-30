// run

package main

type C[T any] struct {
}

func (c *C[T]) reset() {
}

func New[T any]() {
	c := &C[T]{}
	i = c.reset
	z(c.reset)
}

var i interface{}

func z(interface{}) {
}

func main() {
	New[int]()
}
