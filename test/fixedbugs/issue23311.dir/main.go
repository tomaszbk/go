package main

import _ "unsafe" // for linkname

//go:linkname f runtime.GC
func f()

func main() {
	f()
}
