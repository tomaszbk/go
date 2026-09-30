// The scenarios written with existing Go syntax: C functions called in cgo's
// two-result form, and explicit error checks.
package main

/*
#include <stdlib.h>
#include "clib.h"
*/
import "C"

import (
	"fmt"
	"strconv"
	"unsafe"
)

func single(num, den int) (int, error) {
	mark("before")
	r, err := C.c_div(C.int(num), C.int(den))
	if err != nil {
		return 0, err
	}
	mark("after")
	return int(r) + 1, nil
}

func wrapped(num, den int) (int, error) {
	mark("before")
	r, err := C.c_div(C.int(num), C.int(den))
	if err != nil {
		mark("handler")
		return -1, fmt.Errorf("divide %d/%d: %w", num, den, err)
	}
	mark("after")
	return int(r), nil
}

func discard(den int) error {
	_, err := C.c_div(10, C.int(den))
	if err != nil {
		return err
	}
	mark("after")
	return nil
}

func voidStmt(x int) error {
	_, err := C.c_touch(C.int(x))
	if err != nil {
		return err
	}
	mark("after")
	return nil
}

func voidHandled(x int) error {
	_, err := C.c_touch(C.int(x))
	if err != nil {
		C.c_note_handled()
		return fmt.Errorf("touch %d: %w", x, err)
	}
	mark("after")
	return nil
}

func args(d1, d2 int) (int, error) {
	a, err := C.c_div(100, C.int(d1))
	if err != nil {
		return 0, err
	}
	b := side("middle", 10)
	c, err := C.c_div(50, C.int(d2))
	if err != nil {
		return 0, err
	}
	r := C.c_add3(a, C.int(b), c)
	mark("done")
	return int(r), nil
}

func goArg(num, den int) (int, error) {
	r, err := C.c_div(C.int(num), C.int(den))
	if err != nil {
		return 0, err
	}
	return sum2(int(r), side("second", 100)), nil
}

func parse(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return n * 2, nil
}

func parseHandled(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		mark("handler")
		return -1, fmt.Errorf("parse: %w", err)
	}
	return n * 2, nil
}

func pointerCheck(n, k int) (int, error) {
	b := C.c_buf{n: C.int(n)}
	r, err := C.c_bufsum(&b, C.int(k))
	if err != nil {
		return 0, err
	}
	mark("done")
	return int(r), nil
}

func pointerHandled(n, k int) (int, error) {
	b := C.c_buf{n: C.int(n)}
	r, err := C.c_bufsum(&b, C.int(k))
	if err != nil {
		mark("handler")
		return -1, err
	}
	return int(r), nil
}

// The function literal passed to a rewritten call keeps its own error handling.
func literalArg(s string) (int, error) {
	b := C.c_buf{n: 100}
	r, err := C.c_bufsum(&b, C.int(func() int {
		n, err := strconv.Atoi(s)
		if err != nil {
			return -1
		}
		return n
	}()))
	if err != nil {
		return 0, err
	}
	return int(r), nil
}

func literal(num, den int) (int, error) {
	inner := func() (int, error) {
		mark("inner")
		r, err := C.c_div(C.int(num), C.int(den))
		if err != nil {
			return 0, err
		}
		return int(r), nil
	}
	v, err := inner()
	mark("outer")
	if err != nil {
		return -2, err
	}
	return v, nil
}

func named(num, den int) (n int, err error) {
	n = 99
	defer func() {
		mark(fmt.Sprintf("defer:%d:%s", n, describe(err)))
		n += 10
	}()
	r, callErr := C.c_div(C.int(num), C.int(den))
	if callErr != nil {
		return 0, callErr
	}
	n = int(r)
	return
}

func cleanup(s string) (int, error) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	defer mark("cleanup")
	n, err := C.c_strlen(cs)
	if err != nil {
		return 0, err
	}
	mark("measured")
	return int(n), nil
}

func loop(start int) (int, error) {
	total := 0
	for i := start; ; i -= 2 {
		r, err := C.c_div(60, C.int(i))
		if err != nil {
			return 0, err
		}
		if !(r < 40) {
			break
		}
		mark(fmt.Sprint("body:", i))
		total += i
	}
	return total, nil
}

func handlerCallsImpl(den, touch int) (int, error) {
	r, err := C.c_div(1, C.int(den))
	if err != nil {
		C.c_note_handled()
		if _, touchErr := C.c_touch(C.int(touch)); touchErr != nil {
			return 0, touchErr
		}
		return -1, err
	}
	return int(r), nil
}
