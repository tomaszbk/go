// run


// The PkgPath of unexported fields of types defined in package main was incorrectly ""

package main

import (
	"fmt"
	"reflect"
)

type foo struct {
	bar int
}

func main() {
	pkgpath := reflect.ValueOf(foo{}).Type().Field(0).PkgPath
	if pkgpath != "main" {
		fmt.Printf("BUG: incorrect PkgPath: %v", pkgpath)
	}
}
