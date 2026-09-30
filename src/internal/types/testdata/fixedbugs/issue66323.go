package p

import "time"

// This type declaration must not cause problems with
// the type validity checker.

type S[T any] struct {
	a T
	b time.Time
}

var _ S[time.Time]
