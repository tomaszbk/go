package main

/*
//enum test { foo, bar };
*/
import "C"

func main() {
	var a = C.enum_test(1) // ERROR HERE
	_ = a
}
