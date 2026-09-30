//go:build !plan9
// +build !plan9

package main

// void start(void);
import "C"

func init() {
	register("CgoExternalThreadPanic", CgoExternalThreadPanic)
}

func CgoExternalThreadPanic() {
	C.start()
	select {}
}

//export gopanic
func gopanic() {
	panic("BOOM")
}
