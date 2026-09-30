package main

// This program will abort.

/*
#include <stdlib.h>
*/
import "C"

func init() {
	register("Abort", Abort)
}

func Abort() {
	C.abort()
}
