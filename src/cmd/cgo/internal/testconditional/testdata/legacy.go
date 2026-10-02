package main

// #include "clib.h"
import "C"

func values(choice bool) (v, sum, nilsum int) {
	a, b := C.c_buf{n: 10}, C.c_buf{n: 20}
	if mark("condition", choice) {
		v = int(C.c_step(C.int(number("then", 3))))
	} else {
		v = int(C.c_step(C.int(number("else", 5))))
	}
	var selected *C.c_buf
	if choice {
		selected = namedBuffer(buffer("left", &a))
	} else {
		selected = buffer("right", &b)
	}
	sum = int(C.c_sum(selected, C.int(number("tail", 7))))
	var absent *C.c_buf
	if choice {
		absent = nil
	} else {
		absent = nil
	}
	nilsum = int(C.c_sum(absent, 2))
	var deferred *C.c_buf
	if mark("defer", choice) {
		deferred = buffer("deferred-pointer", &a)
	} else {
		deferred = buffer("deferred-pointer", nil)
	}
	defer C.c_sum(deferred, C.int(number("deferred-tail", 2)))
	trace = append(trace, "end")
	return
}

func division(choice bool) (int, error) {
	var n C.int
	var err error
	if choice {
		n, err = C.c_div(9, 0)
	} else {
		n, err = C.c_div(9, 3)
	}
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func handled(choice bool) (int, error) {
	var n C.int
	var err error
	if choice {
		n, err = C.c_div(9, 0)
	} else {
		n, err = C.c_div(9, 3)
	}
	if err != nil {
		return -2, err
	}
	return int(n), nil
}

func pointerCheck(choice bool) {
	n := C.int(1)
	b := C.c_buf{p: &n}
	var selected *C.c_buf
	if choice {
		selected = &b
	} else {
		selected = nil
	}
	C.c_sum(selected, 1)
}
