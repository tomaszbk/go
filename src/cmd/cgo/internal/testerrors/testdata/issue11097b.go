package main

/*
//enum test { foo, bar };
*/
import "C"

func main() {
	p := new(C.enum_test) // ERROR HERE
	_ = p
}
