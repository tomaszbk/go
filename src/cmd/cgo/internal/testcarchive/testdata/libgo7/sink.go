package main

import "C"

var sink []byte

//export GoFunction7
func GoFunction7() {
	sink = make([]byte, 4096)
}

func main() {
}
