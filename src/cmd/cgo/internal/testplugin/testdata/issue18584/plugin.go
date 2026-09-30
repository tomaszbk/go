package main

import "reflect"

type C struct {
}

func F(c *C) *C {
	return nil
}

func G() bool {
	var c *C
	return reflect.TypeOf(F).Out(0) == reflect.TypeOf(c)
}
