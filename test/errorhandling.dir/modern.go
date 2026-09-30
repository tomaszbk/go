package main

import (
	"fmt"
	"unsafe"
)

func single(fail bool) (int, error) {
	v := value("single", 7, fail)! // A comment still terminates the statement.
	return v + 1, nil
}

func multiple(fail bool) (int, string, error) {
	n, s := values(fail)!
	return n, s, nil
}

func errorOnly(fail bool) error {
	action("action", fail)! /* trailing comment */
	mark("after")
	return nil
}

func named(fail bool) (n int, s string, err error) {
	n, s = 99, "kept"
	defer func() {
		mark(fmt.Sprintf("defer:%d:%s:%v", n, s, err))
		mark(fmt.Sprint("recover:", recover()))
		n += 10
	}()
	n = value("named", 8, fail)!
	return
}

func argumentOrder(fail bool) (int, error) {
	return sum(side("left", 1), value("middle", 2, fail)!, side("right", 3)), nil
}

func assignmentOrder(fail bool) (int, error) {
	dst := []int{2}
	dst[side("lhs", 0)] += value("rhs", 9, fail)!
	return dst[0], nil
}

func countedLoop(fail bool) (int, error) {
	i, total := 0, 0
	for next(&i, fail)! {
		mark("body")
		total++
	}
	return total, nil
}

func localWrapped(fail bool) (int, error) {
	v := value("wrapped", 12, fail) or err {
		mark("handler")
		return 0, fmt.Errorf("read: %w", err)
	}
	return v, nil
}

func localAction(fail bool) int {
	action("local action", fail) or problem {
		defer mark("handler defer")
		mark(fmt.Sprint("handled:", problem))
	}
	mark("continued")
	return 25
}

func deferredArguments(fail bool) error {
	defer mark("first defer")
	defer consume(value("defer value", 9, fail)!)
	mark("body")
	return nil
}

func closureBoundary(fail bool) (int, error) {
	v, err := func() (int, error) {
		return value("closure", 14, fail)!, nil
	}()
	mark("outer continued")
	if err != nil {
		return 20, nil
	}
	return v, nil
}

func genericIdentity[T any](v T, fail bool) (T, error) {
	return genericSource(v, fail)!, nil
}

func iteratorPropagation(fail bool) (n int, err error) {
	defer func() {
		mark(fmt.Sprintf("iterator defer:%d:%v", n, err))
		n += 100
	}()
	for i := range sequence {
		n += value(fmt.Sprint("iterator value:", i), i, fail && i == 2)!
	}
	return
}

func nestedPropagation(fail bool) (int, error) {
	return value("outer call", value("inner call", 19, fail)!, false)!, nil
}

func handlerControl(fail bool) (int, error) {
	v := value("control", 16, fail) or err {
		for i := 0; i < 3; i++ {
			if i == 0 {
				continue
			}
			if i == 2 {
				break
			}
			mark("handler iteration")
		}
		return -1, err
	}
	return v, nil
}

func expressionContexts(fail bool) (int, error) {
	m := map[int]int{side("map key", 1): value("map value", 4, fail)!}
	s := []int{value("slice value", 5, false)!}
	ch := make(chan int, 1)
	select {
	case ch <- value("send value", 6, false)!:
	}
	switch 1 {
	case value("switch case", 1, false)!:
	case value("skipped switch case", 2, true)!:
		panic("wrong case")
	default:
		panic("missing case")
	}
	return m[1] + s[0] + <-ch, nil
}

func zeroValues() (*int, []int, map[int]int, chan int, func(), any, error) {
	action("zero values", true)!
	panic("unreachable")
}

func unevaluatedOperand() (uintptr, error) {
	const size = unsafe.Sizeof(value("must not evaluate", 1, true)!)
	return size, nil
}

func logical(first, useOr, fail bool) (bool, error) {
	if useOr {
		return first || boolean("logical", true, fail)!, nil
	}
	return first && boolean("logical", true, fail)!, nil
}

func nilInterfaceSemantics() (int, error) {
	return typedNil()!, nil
}

func aliasedError() (int, errorAlias) {
	return aliasSource()!, nil
}

func compatibility() bool {
	or := false
	a := !or
	b := ! or // Intentionally test same-line whitespace after prefix negation.
	return a && b && or != true
}

func handlerScope() int {
	err := 5
	func() error { return nil }() or err {
		return 0
	}
	return err
}

func handlerPanic() (s string) {
	defer func() { s = fmt.Sprint(recover()) }()
	v := value("panic", 1, true) or err {
		panic(err)
	}
	return fmt.Sprint(v)
}

func normalPanic() (s string) {
	defer func() { s = fmt.Sprint(recover()) }()
	action("before panic", false) or err {
		return err.Error()
	}
	panic("ordinary panic")
}
