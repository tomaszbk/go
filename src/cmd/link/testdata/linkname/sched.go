package main

import _ "unsafe"

type schedt struct{}

//go:linkname sched runtime.sched
var sched schedt

func main() {
	select {
	default:
		println("hello")
	}
}
