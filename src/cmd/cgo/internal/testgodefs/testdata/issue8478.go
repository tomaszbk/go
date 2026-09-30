//go:build ignore

package main

// Issue 8478.  Test that void* is consistently mapped to *byte.

/*
typedef struct {
	void *p;
	void **q;
	void ***r;
} s;
*/
import "C"

type Issue8478 C.s
