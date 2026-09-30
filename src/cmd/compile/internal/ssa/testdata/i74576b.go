package main

import (
	"runtime"
)

func main() {
	a := 1
	runtime.Breakpoint()
	_ = make([]int, a)
}
