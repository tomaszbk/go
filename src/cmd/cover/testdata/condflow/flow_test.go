package flow

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"testing"
)

// Each path runs in a fresh process so that its profile independently records
// selected branches, handler bodies, and code after an early return.
func TestThenOK(t *testing.T)    { testPath(t, true, false) }
func TestElseOK(t *testing.T)    { testPath(t, false, false) }
func TestThenError(t *testing.T) { testPath(t, true, true) }
func TestElseError(t *testing.T) { testPath(t, false, true) }

func testPath(t *testing.T, choose, fail bool) {
	t.Helper()
	checkCalls := func(want ...string) {
		t.Helper()
		if !slices.Equal(calls, want) {
			t.Fatalf("calls = %q, want %q", calls, want)
		}
	}

	calls = nil
	wantSelect := 3
	if choose {
		wantSelect = 2
	}
	for range 2 {
		if got := Select(choose); got != wantSelect {
			t.Fatalf("Select = %d, want %d", got, wantSelect)
		}
	}
	s := strconv.Itoa(wantSelect)
	checkCalls("predicate", s, "selected", "predicate", s, "selected")

	// An invalid value in the unused branch must not cause a failure or a call.
	selected, other := "7", "x"
	if fail {
		selected, other = "x", "7"
	}
	a, b := other, selected
	branch := "else"
	if choose {
		a, b = selected, other
		branch = "then"
	}
	checkResult := func(name string, got, want int, err error, prefix string) {
		t.Helper()
		if got != want {
			t.Fatalf("%s = %d, want %d", name, got, want)
		}
		if !fail {
			if err != nil {
				t.Fatalf("%s: %v", name, err)
			}
			return
		}
		if !errors.Is(err, strconv.ErrSyntax) || !strings.HasPrefix(err.Error(), prefix) {
			t.Fatalf("%s error = %v, want prefix %q wrapping ErrSyntax", name, err, prefix)
		}
	}

	calls = nil
	n, err := Propagate(choose, a, b)
	if fail {
		checkResult("Propagate", n, 0, err, "strconv.Atoi:")
		checkCalls("predicate", selected, "defer:0")
	} else {
		checkResult("Propagate", n, 8, err, "")
		checkCalls("predicate", selected, "propagated", "defer:8")
	}

	calls = nil
	n, err = Handle(choose, a, b)
	if fail {
		checkResult("Handle", n, 0, err, branch+": ")
		checkCalls("predicate", selected, branch+"-handler")
	} else {
		checkResult("Handle", n, 8, err, "")
		checkCalls("predicate", selected, "handled")
	}

	calls = nil
	n, err = Closure(choose, a, b)
	if fail {
		checkResult("Closure", n, -1, err, "strconv.Atoi:")
		checkCalls("predicate", selected)
	} else {
		checkResult("Closure", n, 8, err, "")
		checkCalls("predicate", selected, branch+"-closure", "closed")
	}

	calls = nil
	n = Nested(choose)
	if choose {
		if n != 5 {
			t.Fatalf("Nested = %d, want 5", n)
		}
		checkCalls("predicate", "predicate", "5", "nested")
	} else {
		if n != 6 {
			t.Fatalf("Nested = %d, want 6", n)
		}
		checkCalls("predicate", "6", "nested")
	}
}
