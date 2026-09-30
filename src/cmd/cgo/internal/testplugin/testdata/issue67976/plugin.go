package main

import (
	"io"
	"runtime/pprof"
)

func main() {}

func Start() {
	pprof.StartCPUProfile(io.Discard)
}
