package flow

import (
	"fmt"
	"strconv"
)

func Select(choose bool) int {
	n := if predicate(choose) { value(2) } else { value(3) } // @select
	calls = append(calls, "selected")                        // @select-after
	return n
}

func Propagate(choose bool, a, b string) (n int, err error) {
	defer func() {
		calls = append(calls, "defer:"+strconv.Itoa(n)) // @prop-defer
	}()
	n = if predicate(choose) { atoi(a)! } else { atoi(b)! } // @prop-select
	calls = append(calls, "propagated")                     // @prop-after
	return n + 1, nil
}

func Handle(choose bool, a, b string) (int, error) {
	n := if predicate(choose) { // @handle-select
		atoi(a) or err {
			calls = append(calls, "then-handler") // @handle-then
			return 0, fmt.Errorf("then: %w", err)
		}
	} else {
		atoi(b) or err {
			calls = append(calls, "else-handler") // @handle-else
			return 0, fmt.Errorf("else: %w", err)
		}
	}
	calls = append(calls, "handled") // @handle-after
	return n + 1, nil
}

func Closure(choose bool, a, b string) (int, error) {
	f := if predicate(choose) {
		func() (int, error) {
			n := atoi(a)!                         // @closure-then-parse
			calls = append(calls, "then-closure") // @closure-then-after
			return n + 1, nil
		}
	} else {
		func() (int, error) {
			n := atoi(b)!                         // @closure-else-parse
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
	n := if predicate(choose) { (if predicate(!choose) { value(4) } else { value(5) }) + 0 } else { value(6) } // @nested
	calls = append(calls, "nested")                                                                            // @nested-after
	return n
}
