// compile

package main

type Handler func(in ...interface{})

type Foo[T any] struct{}

func (b *Foo[T]) Bar(in ...interface{}) {}

func (b *Foo[T]) Init() {
	_ = Handler(b.Bar)
}

func main() {
	c := &Foo[int]{}
	c.Init()
}
