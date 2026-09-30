package main

/*
void f(){}
int g(double x){}
*/
import "C"

func init() {
	C.f()
	C.g(0)
}
