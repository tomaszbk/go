package flow

import (
	"errors"
	"strconv"
)

// calls records every call to atoi and check in order, so that tests can
// observe evaluation order and short-circuiting.
var calls []string

func atoi(s string) (int, error) {
	calls = append(calls, s)
	return strconv.Atoi(s)
}

func check(fail bool) error {
	calls = append(calls, "check")
	if fail {
		return errors.New("check failed")
	}
	return nil
}
