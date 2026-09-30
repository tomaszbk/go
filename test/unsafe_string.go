// run


package main

import (
	"unsafe"
)

func main() {
	hello := [5]byte{'m', 'o', 's', 'h', 'i'}
	if unsafe.String(&hello[0], uint64(len(hello))) != "moshi" {
		panic("unsafe.String convert error")
	}
}
