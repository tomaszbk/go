package nilchecks

import "unsafe"

var global *int

func selectors(p *value, fallback *int) {
	if p.P == nil {
		p.P = fallback
	}
	if global == nil {
		global = fallback
	}
}

func nullableField(p *value, fallback *int) *int {
	if p != nil {
		return p.P
	} else {
		return fallback
	}
}

func nullableCall(f func() *int, fallback *int) *int {
	if f != nil {
		return f()
	} else {
		return fallback
	}
}

func boxingField(p *value) any {
	if p != nil {
		return p.P
	} else {
		return nil
	}
}

func comment(p *int, fallback *int) {
	if p == nil {
		// This explanation must not disappear.
		p = fallback
	}
}

func shadowed(p, nil, fallback *int) *int {
	if p == nil {
		p = fallback
	}
	return p
}

func compound(f func()) {
	if f != nil {
		println("effect")
		f()
	}
}

func deferred(f func()) {
	if f != nil {
		defer f()
	}
}

func nonlocal(get func() *value) int {
	if get() != nil {
		return get().N
	} else {
		return 0
	}
}

func incompatibleDefault(p *int) any {
	if p != nil {
		return p
	} else {
		return 42
	}
}

func genericGuard[T ~*int](p, fallback T) T {
	if p == nil {
		p = fallback
	}
	return p
}

func chained(f func() func()) {
	if f != nil {
		f()()
	}
}

func nullableUnsafe(p *struct{ P unsafe.Pointer }, fallback unsafe.Pointer) unsafe.Pointer {
	if p != nil {
		return p.P
	} else {
		return fallback
	}
}
