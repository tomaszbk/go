package flow

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"testing"
)

// The three paths below are run separately, each in a fresh process, so that
// their coverage profiles can be compared. The same tests run against the
// legacy and the modern implementation of the package.

func expect[T comparable](t *testing.T, what string, got, want T, err error, wantErr string) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %v, want %v", what, got, want)
	}
	switch {
	case wantErr == "" && err != nil:
		t.Errorf("%s: unexpected error %v", what, err)
	case wantErr != "" && err == nil:
		t.Errorf("%s: missing error %q", what, wantErr)
	case err != nil && err.Error() != wantErr:
		t.Errorf("%s: error %q, want %q", what, err, wantErr)
	}
	// Propagated and wrapped errors keep their identity.
	if strings.HasSuffix(wantErr, errX) && !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("%s: error %v does not wrap strconv.ErrSyntax", what, err)
	}
}

func expectCalls(t *testing.T, want ...string) {
	t.Helper()
	if !slices.Equal(calls, want) {
		t.Errorf("calls = %q, want %q", calls, want)
	}
}

const (
	errX = `strconv.Atoi: parsing "x": invalid syntax`
)

func TestPathOK(t *testing.T) {
	calls = nil
	sum, err := Sum("2", "3")
	expect(t, "Sum", sum, 5, err, "")
	scaled, err := Scale("4", 3)
	expect(t, "Scale", scaled, 12, err, "")
	var notes []string
	if got := Note(false, &notes); got != 1 || !slices.Equal(notes, []string{"done"}) {
		t.Errorf("Note = %d %q", got, notes)
	}
	sign, err := Classify("-7")
	expect(t, "Classify", sign, "negative", err, "")
	total, err := Total([]string{"1", "2"})
	expect(t, "Total", total, 3, err, "")
	doubled, err := Dispatch("double", "5")
	expect(t, "Dispatch double", doubled, 10, err, "")
	negated, err := Dispatch("negate", "6")
	expect(t, "Dispatch negate", negated, -6, err, "")
	unknown, err := Dispatch("other", "1")
	expect(t, "Dispatch unknown", unknown, 0, err, `unknown kind "other"`)
	applied, err := Apply("9")
	expect(t, "Apply", applied, 10, err, "")
	expectCalls(t, "2", "3", "4", "check", "-7", "1", "2", "5", "6", "9")
}

func TestPathBadFirst(t *testing.T) {
	calls = nil
	sum, err := Sum("x", "3")
	expect(t, "Sum", sum, 0, err, errX)
	scaled, err := Scale("x", 3)
	expect(t, "Scale", scaled, 0, err, `scale "x": `+errX)
	var notes []string
	if got := Note(true, &notes); got != 2 || !slices.Equal(notes, []string{"check failed", "done"}) {
		t.Errorf("Note = %d %q", got, notes)
	}
	sign, err := Classify("x")
	expect(t, "Classify", sign, "", err, errX)
	total, err := Total([]string{"x", "2"})
	expect(t, "Total", total, 0, err, errX)
	doubled, err := Dispatch("double", "x")
	expect(t, "Dispatch double", doubled, 0, err, errX)
	negated, err := Dispatch("negate", "x")
	expect(t, "Dispatch negate", negated, 0, err, `negate "x": `+errX)
	// A failure inside the closure returns from the closure only.
	applied, err := Apply("x")
	expect(t, "Apply", applied, -1, err, errX)
	expectCalls(t, "x", "x", "check", "x", "x", "x", "x", "x")
}

func TestPathBadSecond(t *testing.T) {
	calls = nil
	sum, err := Sum("2", "x")
	expect(t, "Sum", sum, 0, err, errX)
	sign, err := Classify("7")
	expect(t, "Classify", sign, "non-negative", err, "")
	total, err := Total([]string{"1", "x", "3"})
	expect(t, "Total", total, 0, err, errX)
	expectCalls(t, "2", "x", "7", "1", "x")
}
