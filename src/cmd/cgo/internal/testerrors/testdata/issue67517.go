package main

// typedef struct { int a; void* ptr; } S;
// static void f(S* p) {}
import "C"

func main() {
	C.f(&C.S{
		a: 1+

			(3 + ""), // ERROR HERE

		ptr: nil,
	})
}
