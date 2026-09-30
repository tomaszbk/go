// build

//go:build cgo


package main

// #cgo CFLAGS: -Werror -Wunused-parameter
import "C"

func main() {
}

//export Fn
func Fn() {
}
