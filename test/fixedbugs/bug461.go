// run

// part two of issue 4124. Make sure reflect doesn't mark the field as exported.

package main

import "reflect"

var T struct {
	int
}

func main() {
	v := reflect.ValueOf(&T)
	v = v.Elem().Field(0)
	if v.CanSet() {
		panic("int should be unexported")
	}
}
