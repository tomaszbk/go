package main

import (
	"runtime"
)

func main() {
	s := S{1, 1}
	runtime.Breakpoint()
	sink = s
}

type S struct{ a, b uint64 }

var sink any
