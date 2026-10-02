package lib

import "errors"

var Failure = errors.New("export failure")

type Err struct{}

func (*Err) Error() string { return "typed nil" }

func Source(fail bool) (int, error) {
	if fail {
		return 99, Failure
	}
	return 7, nil
}

// Small functions with conditional expressions stay inlinable, including
// in importing packages.
func Pick(c bool, a, b int) int { return if c { a } else { b } }

func Label(n int) string { return if n == 1 { "item" } else { "items" } }

func Lazy(c bool, then, els func() int) int { return if c { then() } else { els() } }

func ErrorOf(failed bool) error {
	var p *Err // nil
	return if failed { p } else { nil }
}

func Value(c, fail bool) (int, error) {
	return if c { Source(fail)! + 1 } else { 0 }, nil
}

func Generic[T any](c bool, a, b T) T { return if c { a } else { b } }

func OrDefault[K comparable, V any](m map[K]V, k K, def V) V {
	v, ok := m[k]
	return if ok { v } else { def }
}

type Box[T any] struct {
	V  T
	OK bool
}

func (b Box[T]) Get(def T) T { return if b.OK { b.V } else { def } }

const Size = if ^uint(0)>>63 == 1 { 64 } else { 32 }

var Mode = if Size == 64 { "wide" } else { "narrow" }
