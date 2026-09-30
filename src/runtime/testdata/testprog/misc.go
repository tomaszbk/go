package main

import "runtime"

func init() {
	register("NumGoroutine", NumGoroutine)
}

func NumGoroutine() {
	println(runtime.NumGoroutine())
}
