package flow

import (
	"fmt"
	"strconv"
)

func Select(choose bool) int {
	var n int
	if predicate(choose) { // @select
		n = value(2)
	} else {
		n = value(3)
	}
	calls = append(calls, "selected") // @select-after
	return n
}

func Propagate(choose bool, a, b string) (n int, err error) {
	defer func() {
		calls = append(calls, "defer:"+strconv.Itoa(n)) // @prop-defer
	}()
	if predicate(choose) { // @prop-select
		n, err = atoi(a)
	} else {
		n, err = atoi(b)
	}
	if err != nil {
		return 0, err
	}
	calls = append(calls, "propagated") // @prop-after
	return n + 1, nil
}

func Handle(choose bool, a, b string) (int, error) {
	var n int
	var err error
	if predicate(choose) { // @handle-select
		n, err = atoi(a)
		if err != nil {
			calls = append(calls, "then-handler") // @handle-then
			return 0, fmt.Errorf("then: %w", err)
		}
	} else {
		n, err = atoi(b)
		if err != nil {
			calls = append(calls, "else-handler") // @handle-else
			return 0, fmt.Errorf("else: %w", err)
		}
	}
	calls = append(calls, "handled") // @handle-after
	return n + 1, nil
}

func Closure(choose bool, a, b string) (int, error) {
	var f func() (int, error)
	if predicate(choose) {
		f = func() (int, error) {
			n, err := atoi(a) // @closure-then-parse
			if err != nil {
				return 0, err
			}
			calls = append(calls, "then-closure") // @closure-then-after
			return n + 1, nil
		}
	} else {
		f = func() (int, error) {
			n, err := atoi(b) // @closure-else-parse
			if err != nil {
				return 0, err
			}
			calls = append(calls, "else-closure") // @closure-else-after
			return n + 1, nil
		}
	}
	n, err := f() // @closure-call
	if err != nil {
		return -1, err // @closure-fail
	}
	calls = append(calls, "closed") // @closure-after
	return n, nil
}

func Nested(choose bool) int {
	var n int
	if predicate(choose) { // @nested
		if predicate(!choose) {
			n = value(4)
		} else {
			n = value(5)
		}
	} else {
		n = value(6)
	}
	calls = append(calls, "nested") // @nested-after
	return n
}
