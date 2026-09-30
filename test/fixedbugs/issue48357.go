// run


package main

import "reflect"

type T [129]byte

func main() {
	m := map[string]T{}
	v := reflect.ValueOf(m)
	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(T{}))
	g = m["a"]
}

var g T
