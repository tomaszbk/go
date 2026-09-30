// Shared by legacy.go and modern.go. It is ordinary Go (plus cgo) so that the
// legacy program also builds with an unmodified toolchain. The functions that
// the cases call are defined once in legacy.go (explicit error checks) and once
// in modern.go (postfix ! and "or err { ... }"); both must satisfy the same
// expectations below, so their observable behavior is identical.
package main

/*
#include "clib.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
)

// trace records labels together with the number of traced C calls completed so
// far, which exposes the order of Go and C side effects.
var trace []string

func mark(label string) { trace = append(trace, fmt.Sprintf("%s@%d", label, C.c_ncalls())) }

func side(label string, n int) int { mark(label); return n }

func sum2(a, b int) int { mark("sum2"); return a + b }

// describe renders an error without depending on the platform's errno texts.
func describe(err error) string {
	if err == nil {
		return "<nil>"
	}
	var errno syscall.Errno
	if !errors.As(err, &errno) {
		return "error: " + err.Error()
	}
	name := errno.Error()
	switch errno {
	case syscall.EDOM:
		name = "EDOM"
	case syscall.EINVAL:
		name = "EINVAL"
	}
	if err == error(errno) {
		return "errno " + name
	}
	return "wraps errno: " + strings.ReplaceAll(err.Error(), errno.Error(), name)
}

func ret(n int, err error) string { return fmt.Sprintf("%d, %s", n, describe(err)) }

func handled() int { return int(C.c_nhandled()) }

// Explicit tuple handling keeps working next to the new syntax, including the
// partial result that propagation would discard.
func explicitTuple(den int) (int, error) {
	r, err := C.c_div(7, C.int(den))
	mark("tuple")
	return int(r), err
}

func explicitVoid(x int) error {
	_, err := C.c_touch(C.int(x))
	mark("void")
	return err
}

type testCase struct {
	name string
	run  func() string
	want string
}

var cases = []testCase{
	{"single/ok", func() string { return ret(single(10, 2)) }, "6, <nil> | trace=before@0,after@1 | c-calls=1"},
	{"single/errno", func() string { return ret(single(1, 0)) }, "0, errno EDOM | trace=before@0 | c-calls=1"},

	{"wrapped/ok", func() string { return ret(wrapped(9, 3)) }, "3, <nil> | trace=before@0,after@1 | c-calls=1"},
	{"wrapped/errno", func() string { return ret(wrapped(1, 0)) }, "-1, wraps errno: divide 1/0: EDOM | trace=before@0,handler@1 | c-calls=1"},

	{"discard/ok", func() string { return describe(discard(1)) }, "<nil> | trace=after@1 | c-calls=1"},
	{"discard/errno", func() string { return describe(discard(0)) }, "errno EDOM | trace= | c-calls=1"},

	{"void/ok", func() string { return describe(voidStmt(1)) }, "<nil> | trace=after@1 | c-calls=1"},
	{"void/errno", func() string { return describe(voidStmt(-1)) }, "errno EINVAL | trace= | c-calls=1"},

	{"voidHandled/ok", func() string { return describe(voidHandled(1)) + fmt.Sprint(" handled=", handled()) }, "<nil> handled=0 | trace=after@1 | c-calls=1"},
	{"voidHandled/errno", func() string { return describe(voidHandled(-1)) + fmt.Sprint(" handled=", handled()) }, "wraps errno: touch -1: EINVAL handled=1 | trace= | c-calls=1"},

	{"args/ok", func() string { return ret(args(10, 5)) }, "30, <nil> | trace=middle@1,done@3 | c-calls=3"},
	{"args/first", func() string { return ret(args(0, 5)) }, "0, errno EDOM | trace= | c-calls=1"},
	{"args/last", func() string { return ret(args(10, 0)) }, "0, errno EDOM | trace=middle@1 | c-calls=2"},

	{"goArg/ok", func() string { return ret(goArg(8, 2)) }, "104, <nil> | trace=second@1,sum2@1 | c-calls=1"},
	{"goArg/errno", func() string { return ret(goArg(8, 0)) }, "0, errno EDOM | trace= | c-calls=1"},

	{"goCall/ok", func() string { return ret(parse("21")) }, "42, <nil> | trace= | c-calls=0"},
	{"goCall/err", func() string { return ret(parse("x")) }, `0, error: strconv.Atoi: parsing "x": invalid syntax | trace= | c-calls=0`},
	{"goCall/handled", func() string { return ret(parseHandled("x")) }, `-1, error: parse: strconv.Atoi: parsing "x": invalid syntax | trace=handler@0 | c-calls=0`},

	{"pointerCheck/ok", func() string { return ret(pointerCheck(5, 2)) }, "7, <nil> | trace=done@1 | c-calls=1"},
	{"pointerCheck/errno", func() string { return ret(pointerCheck(5, 0)) }, "0, errno EDOM | trace= | c-calls=1"},
	{"pointerHandled/errno", func() string { return ret(pointerHandled(5, 0)) }, "-1, errno EDOM | trace=handler@1 | c-calls=1"},

	{"literalArg/ok", func() string { return ret(literalArg("7")) }, "107, <nil> | trace= | c-calls=1"},
	{"literalArg/fallback", func() string { return ret(literalArg("x")) }, "99, <nil> | trace= | c-calls=1"},
	{"literalArg/errno", func() string { return ret(literalArg("0")) }, "0, errno EDOM | trace= | c-calls=1"},

	{"literal/ok", func() string { return ret(literal(6, 3)) }, "2, <nil> | trace=inner@0,outer@1 | c-calls=1"},
	{"literal/errno", func() string { return ret(literal(6, 0)) }, "-2, errno EDOM | trace=inner@0,outer@1 | c-calls=1"},

	{"named/ok", func() string { return ret(named(10, 2)) }, "15, <nil> | trace=defer:5:<nil>@1 | c-calls=1"},
	{"named/errno", func() string { return ret(named(10, 0)) }, "10, errno EDOM | trace=defer:0:errno EDOM@1 | c-calls=1"},

	{"cleanup/ok", func() string { return ret(cleanup("abc")) }, "3, <nil> | trace=measured@1,cleanup@1 | c-calls=1"},
	{"cleanup/errno", func() string { return ret(cleanup("")) }, "0, errno EINVAL | trace=cleanup@1 | c-calls=1"},

	{"loop/ok", func() string { return ret(loop(3)) }, "3, <nil> | trace=body:3@1 | c-calls=2"},
	{"loop/errno", func() string { return ret(loop(4)) }, "0, errno EDOM | trace=body:4@1,body:2@2 | c-calls=3"},

	{"handlerCalls/ok", func() string { return handlerCalls(2, 1) }, "0, <nil> handled=0 | trace= | c-calls=1"},
	{"handlerCalls/errno", func() string { return handlerCalls(0, 1) }, "-1, errno EDOM handled=1 | trace= | c-calls=2"},
	{"handlerCalls/replaced", func() string { return handlerCalls(0, -1) }, "0, errno EINVAL handled=1 | trace= | c-calls=2"},

	{"explicitTuple/ok", func() string { return ret(explicitTuple(7)) }, "1, <nil> | trace=tuple@1 | c-calls=1"},
	{"explicitTuple/partial", func() string { return ret(explicitTuple(0)) }, "-1, errno EDOM | trace=tuple@1 | c-calls=1"},
	{"explicitVoid/errno", func() string { return describe(explicitVoid(-1)) }, "errno EINVAL | trace=void@1 | c-calls=1"},
}

func handlerCalls(den, touch int) string {
	n, err := handlerCallsImpl(den, touch)
	return ret(n, err) + fmt.Sprint(" handled=", handled())
}

func main() {
	failed := false
	for _, c := range cases {
		trace = nil
		C.c_reset()
		result := c.run()
		got := fmt.Sprintf("%s | trace=%s | c-calls=%d", result, strings.Join(trace, ","), C.c_ncalls())
		fmt.Printf("%s: %s\n", c.name, got)
		if got != c.want {
			fmt.Printf("FAIL %s\n got: %s\nwant: %s\n", c.name, got, c.want)
			failed = true
		}
	}
	if failed {
		os.Exit(1)
	}
	fmt.Println("PASS")
}
