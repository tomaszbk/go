package main

/*
char *geterror() {
	return "cgo error";
}
*/
import "C"
import (
	"fmt"
)

func init() {
	register("CgoPanicDeadlock", CgoPanicDeadlock)
}

type cgoError struct{}

func (cgoError) Error() string {
	fmt.Print("") // necessary to trigger the deadlock
	return C.GoString(C.geterror())
}

func CgoPanicDeadlock() {
	panic(cgoError{})
}
