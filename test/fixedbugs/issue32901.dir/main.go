package main

import (
	"reflect"

	"./c"
)

func main() {
	x := c.F()
	p := c.P()
	t := reflect.PointerTo(reflect.TypeOf(x))
	tp := reflect.TypeOf(p)
	if t != tp {
		panic("FAIL")
	}
}
