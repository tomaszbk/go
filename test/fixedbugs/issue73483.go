// run -race

//go:build race && cgo

package main

/*
   int v[8192];
*/
import "C"

var x [8192]C.int

func main() {
	copy(C.v[:], x[:])
}
