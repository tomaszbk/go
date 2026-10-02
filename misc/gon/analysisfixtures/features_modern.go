package main

type node struct {
	Next   *node
	N      int
	Values []int
	F      func(int) int
}

var effects int

func effect(n int) int           { effects = effects*10 + n; return n }
func (p *node) value(n int) *int { v := p.N + n; return &v }

type problem struct{}

func (problem) Error() string { return "failed" }
func read(fail bool) (int, error) {
	if fail {
		return 99, problem{}
	}
	return 7, nil
}
func mapValues[T, U any](xs []T, f func(T) U) []U {
	var out []U
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

var checks int

func check(ok bool) {
	checks++
	if !ok {
		panic(checks)
	}
}

func main() {
	base := 3
	var add func(int) int = (x) => x + base
	base = 4
	check(add(2) == 6)
	xs := mapValues([]int{1, 2}, (x) => x + base)
	check(len(xs) == 2 && xs[0] == 5 && xs[1] == 6)
	var operation func(bool) (int, error) = (fail) => {
		defer func() { effect(9) }()
		v := read(fail)!
		return v, nil
	}
	for _, fail := range []bool{false, true} {
		effects = 0
		v, err := operation(fail)
		check(effects == 9)
		if fail {
			check(v == 0 && err != nil)
		} else {
			check(v == 7 && err == nil)
		}
	}
	var p *node
	effects = 0
	check((p?.Values[effect(1)] ?? effect(2)) == 2 && effects == 2)
	check(p?.value(effect(3)) == nil && effects == 2)
	var callback func(int) int
	check((callback?(effect(4)) ?? 8) == 8 && effects == 2)
	callback = add
	check((callback?(1) ?? 8) == 5)
	p = &node{N: 0, Values: []int{10, 20}}
	check((p?.N ?? 8) == 0)
	check((p?.Next?.N ?? 8) == 8)
	effects = 0
	check((p?.Values[effect(1)] ?? effect(2)) == 20 && effects == 1)
	var pointed *int
	check((*pointed ?? 42) == 42)
	zero := 0
	pointed = &zero
	check((*pointed ?? 42) == 0)
	var boxed any = (*node)(nil) ?? p
	check(boxed == p)
	var held any = (*node)(nil)
	var heldSelected any = held ?? p
	check(heldSelected == held)
	m := map[int]*node{}
	effects = 0
	m[effect(1)] ??= p
	check(effects == 1 && m[1] == p)
	m[1] ??= &node{N: effect(2)}
	check(effects == 1 && m[1] == p)
	extended()
	println("PASS")
}

// These boundaries exercise the intrinsic type of chain results separately
// from their contextual interface conversion and each independent guard.
type nilProblem struct{}

func (*nilProblem) Error() string { return "nil problem" }

type holder struct {
	Err     *nilProblem
	Pointer *int
	Nested  **int
}
type nilReceiver interface{ acceptNil() *node }

func (p *node) acceptNil() *node { effect(5); return p }
func (p *node) failure(fail bool) (*node, error) {
	effect(6)
	if fail {
		return p, problem{}
	}
	return p, nil
}
func (p *node) errorOnly(fail bool) error {
	effect(6)
	if fail {
		return problem{}
	}
	return nil
}

func genericNil[T ~[]int | ~map[int]int](a, b T) T { return a ?? b }
func asInterface[T ~*node](a, b T) any             { return a ?? b }
func propagateChain(p *node, fail bool) (*node, error) {
	defer func() { effect(9) }()
	v := p?.failure(fail)!
	return v, nil
}
func handlerChain(p *node, fail bool) *node {
	return p?.failure(fail) or err {
		effect(7)
		return nil
	}
}
func statementChain(p *node, fail bool) error { p?.errorOnly(fail)!; return nil }
func extended() {
	var h *holder
	var err error = h?.Err
	check(err == nil)
	h = &holder{}
	err = h?.Err
	check(err != nil) // present typed nil gets boxed, absence does not
	var chosen error = h?.Err ?? problem{}
	check(chosen.Error() == "failed") // nil tested before interface conversion
	var absent *holder
	chosen = absent?.Err ?? problem{}
	check(chosen.Error() == "failed")
	var n *node
	var iface nilReceiver = n
	effects = 0
	check(iface?.acceptNil() == nil && effects == 5)
	iface = nil
	effects = 0
	check(iface?.acceptNil() == nil && effects == 0)
	var outer **int
	check((**outer ?? 42) == 42)
	inner := (*int)(nil)
	outer = &inner
	check((**outer ?? 42) == 42)
	value := 0
	inner = &value
	check((**outer ?? 42) == 0)
	check((**absent?.Nested ?? 42) == 42)
	h.Nested = outer
	check((**h?.Nested ?? 42) == 0)
	check(len(genericNil([]int(nil), []int{1})) == 1)
	check(len(genericNil(map[int]int(nil), map[int]int{1: 1})) == 1)
	present := &node{N: 7}
	check(asInterface((*node)(nil), present) == present)
	effects = 0
	v, e := propagateChain(nil, true)
	check(v == nil && e == nil && effects == 9)
	effects = 0
	v, e = propagateChain(present, true)
	check(v == nil && e != nil && effects == 69)
	effects = 0
	v, e = propagateChain(present, false)
	check(v == present && e == nil && effects == 69)
	effects = 0
	check(handlerChain(nil, true) == nil && effects == 0)
	effects = 0
	check(handlerChain(present, true) == nil && effects == 67)
	effects = 0
	check(statementChain(nil, true) == nil && effects == 0)
	effects = 0
	check(statementChain(present, true) != nil && effects == 6)
	var optional func(int) int
	optional = optional ?? (x) => x + 1
	check(optional(2) == 3)
	effects = 0
	var funcNode *node
	funcNode = funcNode ?? funcNode ?? &node{N: effect(1)}
	check(funcNode.N == 1 && effects == 1)
	effects = 0
	var stored *node
	stored ??= present
	stored ??= &node{N: effect(2)}
	check(stored == present && effects == 0)
	type slots struct{ P *node }
	obj := &slots{P: present}
	obj.P ??= &node{N: effect(3)}
	check(obj.P == present && effects == 0)
}
