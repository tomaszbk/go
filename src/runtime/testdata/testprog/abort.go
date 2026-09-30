package main

import _ "unsafe" // for go:linkname

func init() {
	register("Abort", Abort)
}

//go:linkname runtimeAbort runtime.abort
func runtimeAbort()

func Abort() {
	defer func() {
		recover()
		panic("BAD: recovered from abort")
	}()
	runtimeAbort()
	println("BAD: after abort")
}
