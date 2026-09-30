package main

// #cgo nocallback annotations for a C function means it should not callback to Go.
// But it do callback to go in this test, Go should crash here.

/*
#cgo nocallback runCShouldNotCallback

extern void runCShouldNotCallback();
*/
import "C"

import (
	"fmt"
)

func init() {
	register("CgoNoCallback", CgoNoCallback)
}

//export CallbackToGo
func CallbackToGo() {
}

func CgoNoCallback() {
	C.runCShouldNotCallback()
	fmt.Println("OK")
}
