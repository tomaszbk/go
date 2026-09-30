// run -gcflags -l=4

package main

import (
	"fmt"
	"runtime"
)

type frame struct {
	pc   uintptr
	file string
	line int
	ok   bool
}

var (
	skip        int
	globalFrame frame
)

func f() {
	g() // line 23
}

func g() {
	h() // line 27
}

func h() {
	x := &globalFrame
	x.pc, x.file, x.line, x.ok = runtime.Caller(skip) // line 32
}

//go:noinline
func testCaller(skp int) frame {
	skip = skp
	f() // line 38
	frame := globalFrame
	if !frame.ok {
		panic(fmt.Sprintf("skip=%d runtime.Caller failed", skp))
	}
	return frame
}

type wantFrame struct {
	funcName string
	line     int
}

// -1 means don't care
var expected = []wantFrame{
	0: {"main.h", 32},
	1: {"main.g", 27},
	2: {"main.f", 23},
	3: {"main.testCaller", 38},
	4: {"main.main", 64},
	5: {"runtime.main", -1},
	6: {"runtime.goexit", -1},
}

func main() {
	for i := 0; i <= 6; i++ {
		frame := testCaller(i) // line 64
		fn := runtime.FuncForPC(frame.pc)
		if expected[i].line >= 0 && frame.line != expected[i].line {
			panic(fmt.Sprintf("skip=%d expected line %d, got line %d", i, expected[i].line, frame.line))
		}
		if fn.Name() != expected[i].funcName {
			panic(fmt.Sprintf("skip=%d expected function %s, got %s", i, expected[i].funcName, fn.Name()))
		}
	}
}
