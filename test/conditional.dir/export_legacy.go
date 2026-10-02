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

func Pick(c bool, a, b int) int {
	if c {
		return a
	}
	return b
}

func Label(n int) string {
	if n == 1 {
		return "item"
	}
	return "items"
}

func Lazy(c bool, then, els func() int) int {
	if c {
		return then()
	}
	return els()
}

func ErrorOf(failed bool) error {
	var p *Err // nil
	if failed {
		return p
	}
	return nil
}

func Value(c, fail bool) (int, error) {
	if !c {
		return 0, nil
	}
	v, err := Source(fail)
	if err != nil {
		return 0, err
	}
	return v + 1, nil
}

func Generic[T any](c bool, a, b T) T {
	if c {
		return a
	}
	return b
}

func OrDefault[K comparable, V any](m map[K]V, k K, def V) V {
	if v, ok := m[k]; ok {
		return v
	}
	return def
}

type Box[T any] struct {
	V  T
	OK bool
}

func (b Box[T]) Get(def T) T {
	if b.OK {
		return b.V
	}
	return def
}

const Size = 32 << (^uint(0) >> 63)

var Mode = func() string {
	if Size == 64 {
		return "wide"
	}
	return "narrow"
}()
