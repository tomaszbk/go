// build

//go:build cgo


package main

/*
#cgo CFLAGS: -Werror -Wimplicit-function-declaration

#include <stdio.h>

static void CFn(_GoString_ gostr) {
	printf("%.*s\n", (int)(_GoStringLen(gostr)), _GoStringPtr(gostr));
}
*/
import "C"

func main() {
	C.CFn("hello, world")
}

// The bug only occurs if there is an exported function.
//export Fn
func Fn() {
}
