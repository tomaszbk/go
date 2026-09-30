package main

/*
long double x = 0;
*/
import "C"

func main() {
	_ = C.x // ERROR HERE
	_ = C.x
}
