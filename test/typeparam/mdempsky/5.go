// compile

package a

type X[T any] int

func (X[T]) F(T) {}

func x() {
	X[interface{}](0).F(0)
}
