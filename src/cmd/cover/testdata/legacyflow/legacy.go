// Package legacyflow is ordinary Go code without any error handling syntax.
// It must be instrumented exactly as it always was; see
// TestLegacyInstrumentationUnchanged.
package legacyflow

import (
	"errors"
	"strconv"
)

// or, err and ok are perfectly good identifiers.
var or = 1

type ok bool

var handler = func(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func contextual(err error, or int) (int, error) {
	if or != 0 && err != nil {
		return or, err
	}
	or = or + 1
	valid := !(or > 3)
	disabled := false
	valid = !disabled && valid
	if !valid || !disabled {
		return or, nil
	}
	return 0, err
}

func checks(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	n *= 2
	m, err := strconv.Atoi(s + "1")
	if err != nil {
		return n, err
	}
	return n + m, nil
}

func closures(xs []int) (total int, err error) {
	add := func(n int) { total += n }
	for _, x := range xs {
		add(x)
	}
	if f := func() bool { return total > 10 }; f() {
		err = errors.New("large")
	}
	for i := func() int { return 0 }(); i < len(xs); i++ {
		total++
	}
	for _, y := range func() []int { return xs }() {
		total += y
	}
	switch z := func() int { return total }(); z {
	case 0:
		total = 1
	default:
		total = 2
	}
	switch func() int { return total }() {
	case 1:
		total++
	}
	var v any = total
	switch w := v.(type) {
	case int:
		total += w
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("recovered")
		}
	}()
	go func() {
		total++
	}()
	return total, err
}

func flow(n int) int {
	if n > 0 {
		n--
	} else if n < 0 {
		n++
	} else {
		n = 5
	}
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer
			}
			if i == 2 {
				break outer
			}
		}
	}
again:
	n++
	if n < 3 {
		goto again
	}
	switch {
	case n > 10:
		n = 10
		fallthrough
	case n > 5:
		n = 5
	default:
	}
	ch := make(chan int, 1)
	select {
	case ch <- n:
	case m := <-ch:
		n = m
	default:
		n = 0
	}
	if n == 100 {
		panic("too large")
	}
	return n
}

// Multi-line statements, comments and blank lines split blocks.
func spread(a, b int) int {
	x := a +
		b

	// A comment.
	y := x *
		2 /* inline */

	const (
		c1 = 1
		c2 = 2
	)
	return x + y + c1 + c2
}

type counter struct{ n int }

func (c *counter) incr(by int) int {
	c.n += by
	if c.n < 0 {
		c.n = 0
	}
	return c.n
}

func generic[T comparable](xs []T, x T) bool {
	for _, y := range xs {
		if y == x {
			return true
		}
	}
	return false
}

func empty() {}

func emptyBlocks(b bool) {
	if b {
	}
	switch {
	}
	for {
		break
	}
	select {}
}
