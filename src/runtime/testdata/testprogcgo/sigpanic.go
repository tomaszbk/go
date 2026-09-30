package main

// This program will crash.
// We want to test unwinding from sigpanic into C code (without a C symbolizer).

/*
#cgo CFLAGS: -O0

char *pnil;

static int f1(void) {
	*pnil = 0;
	return 0;
}
*/
import "C"

func init() {
	register("TracebackSigpanic", TracebackSigpanic)
}

func TracebackSigpanic() {
	C.f1()
}
