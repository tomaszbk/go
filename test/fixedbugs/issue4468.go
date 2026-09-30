// errorcheck


// Issue 4468: go/defer calls may not be parenthesized.

package p

type T int

func (t *T) F() T {
	return *t
}

type S struct {
	t T
}

func F() {
	go F            // ERROR "must be function call"
	defer F         // ERROR "must be function call"
	go (F)		// ERROR "must be function call|must not be parenthesized"
	defer (F)	// ERROR "must be function call|must not be parenthesized"
	go (F())	// ERROR "must be function call|must not be parenthesized"
	defer (F())	// ERROR "must be function call|must not be parenthesized"
	var s S
	(&s.t).F()
	go (&s.t).F()
	defer (&s.t).F()
}
