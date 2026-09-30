package main

/*
#include <stdlib.h>
#include <stdio.h>

int test(int *a) {
	a[2] = 300;  // BOOM
	return a[2];
}
*/
import "C"

import "fmt"

var cIntArray [2]C.int

func main() {
	r := C.test(&cIntArray[0])
	fmt.Println("r value = ", r)
}
