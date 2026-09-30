package a

import "reflect"

type A = map[int] bool

func F() interface{} {
	return reflect.New(reflect.TypeOf((*A)(nil))).Elem().Interface()
}
