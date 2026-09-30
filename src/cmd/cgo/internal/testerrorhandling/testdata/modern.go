// The same scenarios as legacy.go, written with postfix ! and "or err { ... }".
// Every C call below uses cgo's two-result form implicitly: the error is errno.
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
	r := C.c_div(C.int(num), C.int(den))!
	mark("after")
	return int(r) + 1, nil
}

func wrapped(num, den int) (int, error) {
	mark("before")
	r := C.c_div(C.int(num), C.int(den)) or err {
		mark("handler")
		return -1, fmt.Errorf("divide %d/%d: %w", num, den, err)
	}
	mark("after")
	return int(r), nil
}

func discard(den int) error {
	C.c_div(10, C.int(den))!
	mark("after")
	return nil
}

func voidStmt(x int) error {
	C.c_touch(C.int(x))!
	mark("after")
	return nil
}

func voidHandled(x int) error {
	C.c_touch(C.int(x)) or err {
		C.c_note_handled()
		return fmt.Errorf("touch %d: %w", x, err)
	}
	mark("after")
	return nil
}

func args(d1, d2 int) (int, error) {
	r := C.c_add3(C.c_div(100, C.int(d1))!, C.int(side("middle", 10)), C.c_div(50, C.int(d2))!)
	mark("done")
	return int(r), nil
}

func goArg(num, den int) (int, error) {
	return sum2(int(C.c_div(C.int(num), C.int(den))!), side("second", 100)), nil
}

func parse(s string) (int, error) {
	n := strconv.Atoi(s)!
	return n * 2, nil
}

func parseHandled(s string) (int, error) {
	n := strconv.Atoi(s) or err {
		mark("handler")
		return -1, fmt.Errorf("parse: %w", err)
	}
	return n * 2, nil
}

// cgo rewrites calls that pass pointers to structures holding pointers into a
// function literal that checks them; the error result must still reach the
// enclosing function.
func pointerCheck(n, k int) (int, error) {
	b := C.c_buf{n: C.int(n)}
	r := C.c_bufsum(&b, C.int(k))!
	mark("done")
	return int(r), nil
}

func pointerHandled(n, k int) (int, error) {
	b := C.c_buf{n: C.int(n)}
	r := C.c_bufsum(&b, C.int(k)) or err {
		mark("handler")
		return -1, err
	}
	return int(r), nil
}

// The function literal passed to a rewritten call keeps its own error handling:
// its handler returns from the literal, and the outer ! from the function.
func literalArg(s string) (int, error) {
	b := C.c_buf{n: 100}
	r := C.c_bufsum(&b, C.int(func() int {
		n := strconv.Atoi(s) or err {
			return -1
		}
		return n
	}()))!
	return int(r), nil
}

func literal(num, den int) (int, error) {
	inner := func() (int, error) {
		mark("inner")
		return int(C.c_div(C.int(num), C.int(den))!), nil
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
	r := C.c_div(C.int(num), C.int(den))!
	n = int(r)
	return
}

func cleanup(s string) (int, error) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	defer mark("cleanup")
	n := C.c_strlen(cs)!
	mark("measured")
	return int(n), nil
}

func loop(start int) (int, error) {
	total := 0
	for i := start; C.c_div(60, C.int(i))! < 40; i -= 2 {
		mark(fmt.Sprint("body:", i))
		total += i
	}
	return total, nil
}

func handlerCallsImpl(den, touch int) (int, error) {
	r := C.c_div(1, C.int(den)) or err {
		C.c_note_handled()
		C.c_touch(C.int(touch))!
		return -1, err
	}
	return int(r), nil
}
