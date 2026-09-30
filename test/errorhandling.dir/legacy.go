package main

import (
	"fmt"
	"unsafe"
)

func single(fail bool) (int, error) {
	v, err := value("single", 7, fail)
	if err != nil {
		return 0, err
	}
	return v + 1, nil
}

func multiple(fail bool) (int, string, error) {
	n, s, err := values(fail)
	if err != nil {
		return 0, "", err
	}
	return n, s, nil
}

func errorOnly(fail bool) error {
	if err := action("action", fail); err != nil {
		return err
	}
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
	v, callErr := value("named", 8, fail)
	if callErr != nil {
		return 0, "", callErr
	}
	n = v
	return
}

func argumentOrder(fail bool) (int, error) {
	a := side("left", 1)
	b, err := value("middle", 2, fail)
	if err != nil {
		return 0, err
	}
	return sum(a, b, side("right", 3)), nil
}

func assignmentOrder(fail bool) (int, error) {
	dst := []int{2}
	idx := side("lhs", 0)
	v, err := value("rhs", 9, fail)
	if err != nil {
		return 0, err
	}
	dst[idx] += v
	return dst[0], nil
}

func countedLoop(fail bool) (int, error) {
	i, total := 0, 0
	for {
		ok, err := next(&i, fail)
		if err != nil {
			return 0, err
		}
		if !ok {
			break
		}
		mark("body")
		total++
	}
	return total, nil
}

func localWrapped(fail bool) (int, error) {
	v, err := value("wrapped", 12, fail)
	if err != nil {
		mark("handler")
		return 0, fmt.Errorf("read: %w", err)
	}
	return v, nil
}

func localAction(fail bool) int {
	if err := action("local action", fail); err != nil {
		defer mark("handler defer")
		mark(fmt.Sprint("handled:", err))
	}
	mark("continued")
	return 25
}

func deferredArguments(fail bool) error {
	defer mark("first defer")
	v, err := value("defer value", 9, fail)
	if err != nil {
		return err
	}
	defer consume(v)
	mark("body")
	return nil
}

func closureBoundary(fail bool) (int, error) {
	v, err := func() (int, error) {
		v, err := value("closure", 14, fail)
		if err != nil {
			return 0, err
		}
		return v, nil
	}()
	mark("outer continued")
	if err != nil {
		return 20, nil
	}
	return v, nil
}

func genericIdentity[T any](v T, fail bool) (T, error) {
	v, err := genericSource(v, fail)
	if err != nil {
		var zero T
		return zero, err
	}
	return v, nil
}

func iteratorPropagation(fail bool) (n int, err error) {
	defer func() {
		mark(fmt.Sprintf("iterator defer:%d:%v", n, err))
		n += 100
	}()
	for i := range sequence {
		v, callErr := value(fmt.Sprint("iterator value:", i), i, fail && i == 2)
		if callErr != nil {
			return 0, callErr
		}
		n += v
	}
	return
}

func nestedPropagation(fail bool) (int, error) {
	v, err := value("inner call", 19, fail)
	if err != nil {
		return 0, err
	}
	v, err = value("outer call", v, false)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func handlerControl(fail bool) (int, error) {
	v, err := value("control", 16, fail)
	if err != nil {
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
	k := side("map key", 1)
	v, err := value("map value", 4, fail)
	if err != nil {
		return 0, err
	}
	m := map[int]int{k: v}
	v, err = value("slice value", 5, false)
	if err != nil {
		return 0, err
	}
	s := []int{v}
	ch := make(chan int, 1)
	v, err = value("send value", 6, false)
	if err != nil {
		return 0, err
	}
	select {
	case ch <- v:
	}
	v, err = value("switch case", 1, false)
	if err != nil {
		return 0, err
	}
	if v != 1 {
		panic("wrong case")
	}
	return m[1] + s[0] + <-ch, nil
}

func zeroValues() (*int, []int, map[int]int, chan int, func(), any, error) {
	if err := action("zero values", true); err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	panic("unreachable")
}

func unevaluatedOperand() (uintptr, error) {
	const size = unsafe.Sizeof(side("must not evaluate", 1))
	return size, nil
}

func logical(first, useOr, fail bool) (bool, error) {
	if useOr {
		if first {
			return true, nil
		}
	} else if !first {
		return false, nil
	}
	v, err := boolean("logical", true, fail)
	if err != nil {
		return false, err
	}
	return v, nil
}

func nilInterfaceSemantics() (int, error) {
	v, err := typedNil()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func aliasedError() (int, errorAlias) {
	v, err := aliasSource()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func compatibility() bool {
	or := false
	a := !or
	b := ! or // Intentionally test same-line whitespace after prefix negation.
	return a && b && or != true
}

func handlerScope() int {
	err := 5
	if err := error(nil); err != nil {
		return 0
	}
	return err
}

func handlerPanic() (s string) {
	defer func() { s = fmt.Sprint(recover()) }()
	v, err := value("panic", 1, true)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(v)
}

func normalPanic() (s string) {
	defer func() { s = fmt.Sprint(recover()) }()
	if err := action("before panic", false); err != nil {
		return err.Error()
	}
	panic("ordinary panic")
}
