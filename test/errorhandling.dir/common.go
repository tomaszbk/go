package main

import (
	"errors"
	"fmt"
	"reflect"
)

var failure = errors.New("failure")
var trace []string

func mark(s string) { trace = append(trace, s) }

func value(label string, n int, fail bool) (int, error) {
	mark(label)
	if fail {
		// A failed conventional call can still have a useful partial result.
		return n, failure
	}
	return n, nil
}

func values(fail bool) (int, string, error) {
	mark("values")
	if fail {
		return 41, "partial", failure
	}
	return 42, "complete", nil
}

func action(label string, fail bool) error {
	mark(label)
	if fail {
		return failure
	}
	return nil
}

func boolean(label string, b, fail bool) (bool, error) {
	mark(label)
	if fail {
		return b, failure
	}
	return b, nil
}

func side(label string, n int) int { mark(label); return n }
func sum(a, b, c int) int          { mark("sum"); return a + b + c }
func consume(n int)                { mark(fmt.Sprint("consume:", n)) }

func next(i *int, fail bool) (bool, error) {
	*i++
	return boolean(fmt.Sprint("next:", *i), *i < 3, fail && *i == 2)
}

type nilError struct{}

func (*nilError) Error() string { return "typed nil" }
func typedNil() (int, error)    { return 99, (*nilError)(nil) }

type errorAlias = error

func aliasSource() (int, errorAlias) { return 17, nil }

func genericSource[T any](v T, fail bool) (T, error) {
	if fail {
		return v, failure
	}
	return v, nil
}

func sequence(yield func(int) bool) {
	mark("iterator begin")
	for i := 1; i <= 3; i++ {
		mark(fmt.Sprint("yield:", i))
		if !yield(i) {
			mark("iterator stopped")
			return
		}
	}
	mark("iterator end")
}

func check(label string, got, want any) {
	if !reflect.DeepEqual(got, want) {
		panic(fmt.Sprintf("%s: got %#v, want %#v", label, got, want))
	}
	fmt.Println(label, got)
}

func checkTrace(want ...string) {
	check("trace", trace, want)
	trace = nil
}

func main() {
	for _, fail := range []bool{false, true} {
		v, err := single(fail)
		if fail {
			check("single failure", []any{v, err}, []any{0, failure})
		} else {
			check("single success", []any{v, err}, []any{8, nil})
		}
		checkTrace("single")

		n, s, err := multiple(fail)
		if fail {
			check("multiple failure", []any{n, s, err}, []any{0, "", failure})
		} else {
			check("multiple success", []any{n, s, err}, []any{42, "complete", nil})
		}
		checkTrace("values")

		check("only error", errorOnly(fail), map[bool]error{false: nil, true: failure}[fail])
		if fail {
			checkTrace("action")
		} else {
			checkTrace("action", "after")
		}

		n, s, err = named(fail)
		if fail {
			check("named failure", []any{n, s, err}, []any{10, "", failure})
			checkTrace("named", "defer:0::failure", "recover:<nil>")
		} else {
			check("named success", []any{n, s, err}, []any{18, "kept", nil})
			checkTrace("named", "defer:8:kept:<nil>", "recover:<nil>")
		}

		v, err = argumentOrder(fail)
		if fail {
			check("arguments failure", []any{v, err}, []any{0, failure})
			checkTrace("left", "middle")
		} else {
			check("arguments success", []any{v, err}, []any{6, nil})
			checkTrace("left", "middle", "right", "sum")
		}

		v, err = assignmentOrder(fail)
		if fail {
			check("assignment failure", []any{v, err}, []any{0, failure})
			checkTrace("lhs", "rhs")
		} else {
			check("assignment success", []any{v, err}, []any{11, nil})
			checkTrace("lhs", "rhs")
		}

		v, err = countedLoop(fail)
		if fail {
			check("loop failure", []any{v, err}, []any{0, failure})
			checkTrace("next:1", "body", "next:2")
		} else {
			check("loop success", []any{v, err}, []any{2, nil})
			checkTrace("next:1", "body", "next:2", "body", "next:3")
		}

		v, err = localWrapped(fail)
		if fail {
			check("wrapped failure", v, 0)
			check("wrapped identity", errors.Is(err, failure), true)
			check("wrapped text", err.Error(), "read: failure")
			checkTrace("wrapped", "handler")
		} else {
			check("wrapped success", []any{v, err}, []any{12, nil})
			checkTrace("wrapped")
		}

		check("local action", localAction(fail), 25)
		if fail {
			checkTrace("local action", "handled:failure", "continued", "handler defer")
		} else {
			checkTrace("local action", "continued")
		}

		check("defer arguments", deferredArguments(fail), map[bool]error{false: nil, true: failure}[fail])
		if fail {
			checkTrace("defer value", "first defer")
		} else {
			checkTrace("defer value", "body", "consume:9", "first defer")
		}

		v, err = closureBoundary(fail)
		if fail {
			check("closure failure handled", []any{v, err}, []any{20, nil})
		} else {
			check("closure success", []any{v, err}, []any{14, nil})
		}
		checkTrace("closure", "outer continued")

		v, err = genericIdentity(31, fail)
		if fail {
			check("generic failure", []any{v, err}, []any{0, failure})
		} else {
			check("generic success", []any{v, err}, []any{31, nil})
		}

		v, err = iteratorPropagation(fail)
		if fail {
			check("iterator failure", []any{v, err}, []any{100, failure})
			checkTrace("iterator begin", "yield:1", "iterator value:1", "yield:2", "iterator value:2", "iterator stopped", "iterator defer:0:failure")
		} else {
			check("iterator success", []any{v, err}, []any{106, nil})
			checkTrace("iterator begin", "yield:1", "iterator value:1", "yield:2", "iterator value:2", "yield:3", "iterator value:3", "iterator end", "iterator defer:6:<nil>")
		}

		v, err = nestedPropagation(fail)
		if fail {
			check("nested failure", []any{v, err}, []any{0, failure})
			checkTrace("inner call")
		} else {
			check("nested success", []any{v, err}, []any{19, nil})
			checkTrace("inner call", "outer call")
		}

		v, err = handlerControl(fail)
		if fail {
			check("handler control failure", []any{v, err}, []any{-1, failure})
			checkTrace("control", "handler iteration")
		} else {
			check("handler control success", []any{v, err}, []any{16, nil})
			checkTrace("control")
		}

		v, err = expressionContexts(fail)
		if fail {
			check("contexts failure", []any{v, err}, []any{0, failure})
			checkTrace("map key", "map value")
		} else {
			check("contexts success", []any{v, err}, []any{15, nil})
			checkTrace("map key", "map value", "slice value", "send value", "switch case")
		}
	}

	for _, useOr := range []bool{false, true} {
		for _, first := range []bool{false, true} {
			for _, fail := range []bool{false, true} {
				v, err := logical(first, useOr, fail)
				skipped := (!useOr && !first) || (useOr && first)
				if skipped {
					check("logical skip", []any{v, err}, []any{first, nil})
					checkTrace()
				} else if fail {
					check("logical failure", []any{v, err}, []any{false, failure})
					checkTrace("logical")
				} else {
					check("logical value", []any{v, err}, []any{true, nil})
					checkTrace("logical")
				}
			}
		}
	}

	n, err := nilInterfaceSemantics()
	check("typed nil value", n, 0)
	check("typed nil is error", err != nil, true)
	check("typed nil concrete", reflect.TypeOf(err).String(), "*main.nilError")

	// Explicit legacy handling retains partial results in either program.
	n, s, err := values(true)
	check("explicit partial results", []any{n, s, err}, []any{41, "partial", failure})
	checkTrace("values")

	n, err = aliasedError()
	check("error alias", []any{n, err}, []any{17, nil})
	check("compatibility", compatibility(), true)
	check("handler scope", handlerScope(), 5)
	check("handler panic", handlerPanic(), "failure")
	checkTrace("panic")
	check("normal panic", normalPanic(), "ordinary panic")
	checkTrace("before panic")
	p, slice, m, ch, fn, iface, err := zeroValues()
	check("nil zero results", []any{p == nil, slice == nil, m == nil, ch == nil, fn == nil, iface == nil, err}, []any{true, true, true, true, true, true, failure})
	checkTrace("zero values")
	zeroSlice, err := genericIdentity([]int{1, 2}, true)
	check("generic nil zero", []any{zeroSlice == nil, err}, []any{true, failure})
	size, err := unevaluatedOperand()
	check("constant sizeof", []any{size > 0, err}, []any{true, nil})
	checkTrace()
}
