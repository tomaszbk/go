// run

package main

type G[T any] interface {
	g()
}

type Foo[T any] struct {
}

func (foo *Foo[T]) g() {

}

func f[T any]() {
	v := []G[T]{}
	v = append(v, &Foo[T]{})
}
func main() {
	f[int]()
}
