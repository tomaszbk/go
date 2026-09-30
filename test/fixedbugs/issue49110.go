// run


package main

import "reflect"

func main() {
	_ = reflect.StructOf([]reflect.StructField{
		{Name: "_", PkgPath: "main", Type: reflect.TypeOf(int(0))},
		{Name: "_", PkgPath: "main", Type: reflect.TypeOf(int(0))},
	})
}
