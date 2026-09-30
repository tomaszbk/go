// run

package main

import "reflect"

type W struct {
	E struct{}
	X *byte
}

func main() {
	w := reflect.ValueOf(W{})
	_ = w.Field(0).Interface()
}
