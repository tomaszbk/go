// Conditional expressions in many contexts, for every analyzer to traverse.
// None of them is reported.

package condexpr

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

const intSize = if ^uint(0)>>63 == 1 { 64 } else { 32 }

var table = [if intSize == 64 { 2 } else { 1 }]int{}

var (
	verbose = len(os.Args) > 5
	mode    = if verbose { "verbose" } else { "quiet" }
)

type point struct{ x, y int }

type name struct{ s string }

func (n name) String() string { return n.s }

func pick[T any](c bool, a, b T) T { return if c { a } else { b } }

func smokeValues(n int, s string, w io.Writer, err error, fs []func() int) (int, error) {
	label := if n == 1 { "item" } else { "items" }
	fmt.Printf("%d %s %s\n", n, label, mode)
	fmt.Println(if n > 0 { "positive" } else { "non-positive" }, len(table))
	fmt.Fprintf(w, if intSize == 64 { "%d-bit\n" } else { "%d-bit (small)\n" }, intSize)
	total := n + if len(s) > 3 { 10 } else { 0 }
	p := if n > 0 { point{1, 2} } else { point{x: 3} }
	var out io.Writer = if w != nil { w } else { os.Stdout }
	var e error = if err != nil { fmt.Errorf("wrap: %w", err) } else { nil }
	var pathErr *os.PathError
	if errors.As(e, &pathErr) || errors.Is(e, io.EOF) {
		return 0, e
	}
	m := map[string]int{label: if total > 0 { total } else { -total }}
	list := []int{if n > 0 { n } else { 0 }, total, p.x, m[label]}
	f := if len(fs) > 0 { fs[0] } else { func() int { return len(list) } }
	str := if n > 0 { name{"a"}.String } else { name{"b"}.String }
	fmt.Fprintln(out, str(), pick(n > 0, "yes", "no"), pick(n > 0, 1.5, 2))
	trimmed := if strings.HasPrefix(s, "x") { strings.TrimPrefix(s, "x") } else { s }
	var counter atomic.Int64
	counter.Store(int64(if n > 0 { n } else { 0 }))
	return if total > 0 { total + len(trimmed) } else { f() }, if e != nil { e } else { nil }
}

func smokeControl(n int, s string, ch chan int) (int, error) {
	parsed := if s != "" { strconv.Atoi(s)! } else { 0 }
	fallback := if n > 0 {
		strconv.Atoi(s) or err {
			return 0, fmt.Errorf("parse %q: %w", s, err)
		}
	} else {
		-1
	}
	ch <- if n > 0 { n } else { 0 }
	go func(v int) { ch <- v }(if n > 1 { n } else { 1 })
	defer fmt.Println(if n > 0 { "done" } else { "skipped" })
	sum := 0
	for i := range (if n > 0 { n } else { 0 }) {
		sum += if i%2 == 0 { i } else { -i }
	}
	switch (if n > 0 { "positive" } else { "other" }) {
	case "positive":
		sum++
	}
	if (if n > 0 { sum > 0 } else { false }) {
		sum--
	}
	return sum + parsed + fallback, nil
}
