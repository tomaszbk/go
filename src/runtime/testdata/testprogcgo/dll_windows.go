package main

/*
#include <windows.h>

DWORD getthread() {
	return GetCurrentThreadId();
}
*/
import "C"
import "runtime/testdata/testprogcgo/windows"

func init() {
	register("CgoDLLImportsMain", CgoDLLImportsMain)
}

func CgoDLLImportsMain() {
	C.getthread()
	windows.GetThread()
	println("OK")
}
