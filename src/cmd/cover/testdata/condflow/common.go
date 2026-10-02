package flow

import "strconv"

var calls []string

func predicate(choose bool) bool {
	calls = append(calls, "predicate")
	return choose
}

func value(n int) int {
	calls = append(calls, strconv.Itoa(n))
	return n
}

func atoi(s string) (int, error) {
	calls = append(calls, s)
	return strconv.Atoi(s)
}
