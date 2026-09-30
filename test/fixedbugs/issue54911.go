// compile

package main

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add() Set[T] {
	return s
}

func (s Set[T]) Copy() Set[T] {
	return Set[T].Add(s)
}

func main() {
	_ = Set[int]{42: {}}
}
