// run

package main

import "unsafe"

//go:noinline
func f(x []byte) bool {
	return unsafe.SliceData(x) != nil
}

//go:noinline
func g(x string) bool {
	return unsafe.StringData(x) != nil
}

func main() {
	if f(nil) {
		panic("bad f")
	}
	if g("") {
		panic("bad g")
	}
}
