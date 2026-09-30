// run


package main

import (
	"fmt"
	"reflect"
)

type Stream[T any] struct {
}

func (s Stream[T]) DropWhile() Stream[T] {
	return Pipe[T, T](s)
}

func Pipe[T, R any](s Stream[T]) Stream[R] {
	it := func(fn func(R) bool) {
	}
	fmt.Println(reflect.TypeOf(it).String())
	return Stream[R]{}
}

func main() {
	s := Stream[int]{}
	s = s.DropWhile()
}
