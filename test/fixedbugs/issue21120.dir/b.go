package b

import "reflect"

type X int

func F1() string {
	type x X

	s := struct {
		*x
	}{nil}
	v := reflect.TypeOf(s)
	return v.Field(0).PkgPath
}

func F2() string {
	type y X

	s := struct {
		*y
	}{nil}
	v := reflect.TypeOf(s)
	return v.Field(0).PkgPath
}
