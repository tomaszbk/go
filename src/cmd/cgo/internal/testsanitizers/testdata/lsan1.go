package main

/*
#include <stdlib.h>

int* test() {
  return malloc(sizeof(int));
}

void clearStack(int n) {
  if (n > 0) {
    clearStack(n - 1);
  }
}

*/
import "C"

//go:noinline
func F() {
	C.test()
}

func clearStack(n int) {
	if n > 0 {
		clearStack(n - 1)
	}
}

func main() {
	// Test should fail: memory allocated by C is leaked.
	F()
	clearStack(100)
	C.clearStack(100)
}
