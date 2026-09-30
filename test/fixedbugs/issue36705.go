// run fake-arg-to-force-use-of-go-run

//go:build cgo && !windows


package main

// #include <stdlib.h>
// #include <unistd.h>
import "C"

import "os"

func main() {
	os.Setenv("FOO", "bar")
	s := C.GoString(C.getenv(C.CString("FOO")))
	if s != "bar" {
		panic("bad setenv, environment variable only has value \"" + s + "\"")
	}
	os.Unsetenv("FOO")
	s = C.GoString(C.getenv(C.CString("FOO")))
	if s != "" {
		panic("bad unsetenv, environment variable still has value \"" + s + "\"")
	}
}
