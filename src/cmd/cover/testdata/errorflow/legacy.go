package flow

import "fmt"

// Sum adds two numbers, stopping at the first one that does not parse.
func Sum(a, b string) (int, error) {
	x, err := atoi(a) // @sum-a
	if err != nil {
		return 0, err // @sum-a-fail
	}
	y, err := atoi(b) // @sum-b
	if err != nil {
		return 0, err // @sum-b-fail
	}
	total := x + y    // @sum-total
	return total, nil // @sum-return
}

// Scale parses a number and multiplies it, adding context to a parse error.
func Scale(a string, factor int) (int, error) {
	n, err := atoi(a) // @scale-parse
	if err != nil {
		return 0, fmt.Errorf("scale %q: %w", a, err) // @scale-handler
	}
	n *= factor   // @scale-mul
	return n, nil // @scale-return
}

// Note records a failure of an error-only call and carries on.
func Note(fail bool, notes *[]string) int {
	if err := check(fail); err != nil { // @note-check
		*notes = append(*notes, err.Error()) // @note-handler
	}
	*notes = append(*notes, "done") // @note-after
	return len(*notes)              // @note-return
}

// Classify reports the sign of a number.
func Classify(s string) (string, error) {
	n, err := atoi(s) // @classify-parse
	if err != nil {
		return "", err // @classify-fail
	}
	if n < 0 { // @classify-test
		return "negative", nil // @classify-neg
	}
	return "non-negative", nil // @classify-nonneg
}

// Total sums numbers, stopping at the first one that does not parse.
func Total(items []string) (int, error) {
	total := 0                   // @total-init
	for _, item := range items { // @total-loop
		n, err := atoi(item) // @total-parse
		if err != nil {
			return 0, err // @total-fail
		}
		total += n // @total-add
	}
	return total, nil // @total-return
}

// Dispatch handles a command in a switch, with a different style per case.
func Dispatch(kind, arg string) (int, error) {
	switch kind {
	case "double":
		n, err := atoi(arg) // @dispatch-double
		if err != nil {
			return 0, err // @dispatch-double-fail
		}
		return n * 2, nil // @dispatch-double-return
	case "negate":
		n, err := atoi(arg) // @dispatch-negate
		if err != nil {
			return 0, fmt.Errorf("negate %q: %w", arg, err) // @dispatch-negate-handler
		}
		return -n, nil // @dispatch-negate-return
	}
	return 0, fmt.Errorf("unknown kind %q", kind) // @dispatch-unknown
}

// Apply handles an error inside a closure; the closure's failure must not
// return from Apply itself.
func Apply(s string) (int, error) {
	parse := func() (int, error) { // @apply-closure
		n, err := atoi(s) // @apply-closure-parse
		if err != nil {
			return 0, err // @apply-closure-fail
		}
		return n + 1, nil // @apply-closure-return
	}
	n, err := parse() // @apply-call
	if err != nil {
		return -1, err // @apply-fail
	}
	return n, nil // @apply-return
}
