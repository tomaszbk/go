// compile

package main

func f1() (x int, y float64) {
	return
}

func f2(x int, y float64) {
	return
}

func main() {
	f2(f1()) // this should be a legal call
}

/*
bug080.go:12: illegal types for operand: CALL
	(<int32>INT32)
	({<x><int32>INT32;<y><float32>FLOAT32;})
*/
