package nilchecks

type value struct {
	N int
	P *int
}

func (v *value) observe(n int) {}

func initialize(p *int, makeValue func() *int) *int {
	// want +1 "nil check can use Gon nil-safety operators"
	if p == nil {
		p = makeValue()
	}
	return p
}

func coalesce(p, fallback *int) *int {
	// want +1 "nil check can use Gon nil-safety operators"
	if p != nil {
		return p
	} else {
		return fallback
	}
}

func reversed(p *int, fallback func() *int) *int {
	// want +1 "nil check can use Gon nil-safety operators"
	if nil == p {
		return fallback()
	} else {
		return p
	}
}

func boxed(p *int) any {
	// want +1 "nil check can use Gon nil-safety operators"
	if p != nil {
		return p
	} else {
		return nil
	}
}

func interfaceCoalesce(p, fallback any) any {
	// want +1 "nil check can use Gon nil-safety operators"
	if p != nil {
		return p
	} else {
		return fallback
	}
}

func access(p *value) int {
	// want +1 "nil check can use Gon nil-safety operators"
	if p != nil {
		return p.N
	} else {
		return 1 + 2
	}
}

func pointerAccess(p *value) *int {
	// want +1 "nil check can use Gon nil-safety operators"
	if p != nil {
		return p.P
	} else {
		return nil
	}
}

func dereference(p *int) int {
	// want +1 "nil check can use Gon nil-safety operators"
	if p == nil {
		return 7
	} else {
		return *p
	}
}

func callback(f func(int) int, arg func() int) int {
	// want +1 "nil check can use Gon nil-safety operators"
	if f != nil {
		return f(arg())
	} else {
		return 8
	}
}

func calls(f func(...int), p *value, arg func() int, args []int) {
	// want +1 "nil check can use Gon nil-safety operators"
	if f != nil {
		f(arg())
	}
	// want +1 "nil check can use Gon nil-safety operators"
	if f != nil {
		f(args...)
	}
	// want +1 "nil check can use Gon nil-safety operators"
	if p != nil {
		p.observe(arg())
	}
}

func assign(p *value) (n int) {
	// want +1 "nil check can use Gon nil-safety operators"
	if p != nil {
		n = p.N
	} else {
		n = 42
	}
	return
}
