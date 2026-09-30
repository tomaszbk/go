// run


package main

import (
	"reflect"
)

func main() {
	pi := new(interface{})
	v := reflect.ValueOf(pi).Elem()
	if v.Kind() != reflect.Interface {
		panic(0)
	}
	if (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && v.IsNil() {
		return
	}
	panic(1)
}
