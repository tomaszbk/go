// run -gcflags=all=-d=checkptr

//go:build cgo


package main

// #include <stdlib.h>
import "C"

func main() {
	C.malloc(100)
}
