// errorcheck -std


package main

func main() {
	println([1]byte{}) // ERROR "illegal types for operand: println"
}
