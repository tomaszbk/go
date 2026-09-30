// run

package main

import (
	"log"
	"reflect"
	"runtime"
)

func hello() string {
	return "Hello World" // line 12
}

func foo() string { // line 15
	x := hello() // line 16
	y := hello() // line 17
	return x + y // line 18
}

func bar() string {
	x := hello() // line 22
	return x
}

// funcPC returns the PC for the func value f.
func funcPC(f interface{}) uintptr {
	return reflect.ValueOf(f).Pointer()
}

// Test for issue #15453. Previously, line 22 would appear in foo().
func main() {
	pc := funcPC(foo)
	f := runtime.FuncForPC(pc)
	for ; runtime.FuncForPC(pc) == f; pc++ {
		file, line := f.FileLine(pc)
		if line == 0 {
			continue
		}
		// Line 12 can appear inside foo() because PC-line table has
		// innermost line numbers after inlining.
		if line != 12 && !(line >= 15 && line <= 18) {
			log.Fatalf("unexpected line at PC=%d: %s:%d\n", pc, file, line)
		}
	}
}
