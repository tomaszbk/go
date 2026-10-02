package main

// #include "clib.h"
import "C"

func values(choice bool) (v, sum, nilsum int) {
	a, b := C.c_buf{n: 10}, C.c_buf{n: 20}
	v = int(if mark("condition", choice) { C.c_step(C.int(number("then", 3))) } else { C.c_step(C.int(number("else", 5))) })
	sum = int(C.c_sum((if choice { namedBuffer(buffer("left", &a)) } else { buffer("right", &b) }), C.int(number("tail", 7))))
	nilsum = int(C.c_sum(if choice { nil } else { nil }, 2))
	defer C.c_sum(if mark("defer", choice) { buffer("deferred-pointer", &a) } else { buffer("deferred-pointer", nil) }, C.int(number("deferred-tail", 2)))
	trace = append(trace, "end")
	return
}

func division(choice bool) (int, error) {
	n := if choice { C.c_div(9, 0)! } else { C.c_div(9, 3)! }
	return int(n), nil
}

func handled(choice bool) (int, error) {
	n := if choice {
		C.c_div(9, 0) or err {
			return -2, err
		}
	} else {
		C.c_div(9, 3) or err {
			return -2, err
		}
	}
	return int(n), nil
}

func pointerCheck(choice bool) {
	n := C.int(1)
	b := C.c_buf{p: &n}
	C.c_sum(if choice { &b } else { nil }, 1)
}
