// run

// Make sure that the line number is reported correctly
// for faulting instructions.

package main

import (
	"fmt"
	"runtime"
)

var x byte
var p *byte

//go:noinline
func f() {
	q := p
	x = 11  // line 19
	*q = 12 // line 20
}
func main() {
	defer func() {
		recover()
		var pcs [10]uintptr
		n := runtime.Callers(1, pcs[:])
		frames := runtime.CallersFrames(pcs[:n])
		for {
			f, more := frames.Next()
			if f.Function == "main.f" && f.Line != 20 {
				panic(fmt.Errorf("expected line 20, got line %d", f.Line))
			}
			if !more {
				break
			}
		}
	}()
	f()
}
