package main

/*
#include <stdlib.h>
#include <stdio.h>

struct ss {
	int *p;
	int len;
	int cap;
};

int test(struct ss *a) {
	struct ss *t = a + 1;
	t->len = 100;          // BOOM
	return t->len;
}
*/
import "C"
import "fmt"

var tt C.struct_ss

func main() {
	r := C.test(&tt)
	fmt.Println("r value = ", r)
}
