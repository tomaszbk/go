package main

import (
	"io"
	"runtime/pprof"
)

import "C"

//export go_start_profile
func go_start_profile() {
	pprof.StartCPUProfile(io.Discard)
}

//export go_stop_profile
func go_stop_profile() {
	pprof.StopCPUProfile()
}

func main() {
}
