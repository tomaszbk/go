package workflow

import "reflect"

type Value interface {
	value() reflect.Value
}

func Const[T any](value T) Value {
	return &constant[T]{value}
}

type constant[T any] struct {
	v T
}

func (c *constant[T]) value() reflect.Value {
	return reflect.ValueOf(c.v)
}
