package main

/*
int f();
int g(int x);
*/
import "C"

func main() {
	C.f()
	C.g(0)
}
