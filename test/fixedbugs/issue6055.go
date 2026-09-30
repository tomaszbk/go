// run

package main

import "runtime"

type Closer interface {
	Close()
}

func nilInterfaceDeferCall() {
	defer func() {
		// make sure a traceback happens with jmpdefer on the stack
		runtime.GC()
	}()
	var x Closer
	defer x.Close()
}

func shouldPanic(f func()) {
	defer func() {
		if recover() == nil {
			panic("did not panic")
		}
	}()
	f()
}

func main() {
	shouldPanic(nilInterfaceDeferCall)
}
