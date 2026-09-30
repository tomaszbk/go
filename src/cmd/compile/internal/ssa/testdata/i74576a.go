package main

import (
	"runtime"
)

func main() {
	a := 1
	runtime.Breakpoint()
	sink = a
}

var sink any
