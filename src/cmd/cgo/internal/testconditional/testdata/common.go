package main

// #include "clib.h"
import "C"

import (
	"fmt"
	"strings"
)

var trace []string

type namedBuffer *C.c_buf

func mark(s string, b bool) bool           { trace = append(trace, s); return b }
func number(s string, n int) int           { trace = append(trace, s); return n }
func buffer(s string, p *C.c_buf) *C.c_buf { trace = append(trace, s); return p }
func must(ok bool, message string) {
	if !ok {
		panic(message)
	}
}

func main() {
	for _, choice := range []bool{true, false} {
		trace = nil
		C.c_reset()
		v, sum, nilsum := values(choice)
		wantV, wantSum, branch, pointer := 5, 27, "else", "right"
		if choice {
			wantV, wantSum, branch, pointer = 3, 17, "then", "left"
		}
		must(v == wantV && sum == wantSum && nilsum == 2, "C results")
		wantTrace := "condition," + branch + "," + pointer + ",tail,defer,deferred-pointer,deferred-tail,end"
		must(strings.Join(trace, ",") == wantTrace, "evaluation order: "+strings.Join(trace, ","))
		must(C.c_calls() == 4, "C call count")
		fmt.Printf("values %v %d %d %d %s calls=%d\n", choice, v, sum, nilsum, strings.Join(trace, ","), C.c_calls())

		C.c_reset()
		n, err := division(choice)
		if choice {
			must(n == 0 && err != nil, "propagation failure")
		} else {
			must(n == 3 && err == nil, "propagation success")
		}
		must(C.c_calls() == 1, "lazy propagation")
		fmt.Printf("division %v %d error=%v calls=%d\n", choice, n, err != nil, C.c_calls())

		C.c_reset()
		n, err = handled(choice)
		if choice {
			must(n == -2 && err != nil, "handler failure")
		} else {
			must(n == 3 && err == nil, "handler success")
		}
		must(C.c_calls() == 1, "lazy handler")
		fmt.Printf("handled %v %d error=%v calls=%d\n", choice, n, err != nil, C.c_calls())
	}
	// A selected Go pointer must still be rejected by cgo's pointer check;
	// selecting nil must not evaluate or check the unused pointer branch.
	for _, choice := range []bool{false, true} {
		C.c_reset()
		panicked := false
		func() { defer func() { panicked = recover() != nil }(); pointerCheck(choice) }()
		must(panicked == choice, "pointer checking")
		wantCalls := C.int(1)
		if choice {
			wantCalls = 0
		}
		must(C.c_calls() == wantCalls, "pointer check before C call")
		fmt.Printf("pointer %v panic=%v calls=%d\n", choice, panicked, C.c_calls())
	}
	fmt.Println("PASS")
}
