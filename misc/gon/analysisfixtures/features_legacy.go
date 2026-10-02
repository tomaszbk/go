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
	var add func(int) int = func(x int) int { return x + base }
	base = 4
	check(add(2) == 6)
	xs := mapValues([]int{1, 2}, func(x int) int { return x + base })
	check(len(xs) == 2 && xs[0] == 5 && xs[1] == 6)
	var operation func(bool) (int, error) = func(fail bool) (int, error) {
		defer func() { effect(9) }()
		v, err := read(fail)
		if err != nil {
			return 0, err
		}
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
	var v int
	if p != nil {
		v = p.Values[effect(1)]
	} else {
		v = effect(2)
	}
	check(v == 2 && effects == 2)
	var result *int
	if p != nil {
		result = p.value(effect(3))
	}
	check(result == nil && effects == 2)
	var callback func(int) int
	v = 8
	if callback != nil {
		v = callback(effect(4))
	}
	check(v == 8 && effects == 2)
	callback = add
	v = 8
	if callback != nil {
		v = callback(1)
	}
	check(v == 5)
	p = &node{N: 0, Values: []int{10, 20}}
	v = 8
	if p != nil {
		v = p.N
	}
	check(v == 0)
	v = 8
	if p != nil && p.Next != nil {
		v = p.Next.N
	}
	check(v == 8)
	effects = 0
	if p != nil {
		v = p.Values[effect(1)]
	} else {
		v = effect(2)
	}
	check(v == 20 && effects == 1)
	var pointed *int
	v = 42
	if pointed != nil {
		v = *pointed
	}
	check(v == 42)
	zero := 0
	pointed = &zero
	v = 42
	if pointed != nil {
		v = *pointed
	}
	check(v == 0)
	var boxed any
	var missing *node
	if missing != nil {
		boxed = missing
	} else {
		boxed = p
	}
	check(boxed == p)
	var held any = (*node)(nil)
	var other any = held
	if other == nil {
		other = p
	}
	check(other == held)
	m := map[int]*node{}
	effects = 0
	key := effect(1)
	if m[key] == nil {
		m[key] = p
	}
	check(effects == 1 && m[1] == p)
	if m[1] == nil {
		m[1] = &node{N: effect(2)}
	}
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

func genericNil[T ~[]int | ~map[int]int](a, b T) T {
	if a != nil {
		return a
	}
	return b
}
func asInterface[T ~*node](a, b T) any {
	if a != nil {
		return a
	}
	return b
}
func propagateChain(p *node, fail bool) (*node, error) {
	defer func() { effect(9) }()
	var v *node
	if p != nil {
		var err error
		v, err = p.failure(fail)
		if err != nil {
			return nil, err
		}
	}
	return v, nil
}
func handlerChain(p *node, fail bool) *node {
	var v *node
	if p != nil {
		var err error
		v, err = p.failure(fail)
		if err != nil {
			effect(7)
			return nil
		}
	}
	return v
}
func statementChain(p *node, fail bool) error {
	if p != nil {
		err := p.errorOnly(fail)
		if err != nil {
			return err
		}
	}
	return nil
}
func extended() {
	var h *holder
	var err error
	if h != nil {
		err = h.Err
	}
	check(err == nil)
	h = &holder{}
	err = nil
	if h != nil {
		err = h.Err
	}
	check(err != nil)
	var chosen error
	if h != nil && h.Err != nil {
		chosen = h.Err
	} else {
		chosen = problem{}
	}
	check(chosen.Error() == "failed")
	var absent *holder
	if absent != nil && absent.Err != nil {
		chosen = absent.Err
	} else {
		chosen = problem{}
	}
	check(chosen.Error() == "failed")
	var n *node
	var iface nilReceiver = n
	var receiver *node
	effects = 0
	if iface != nil {
		receiver = iface.acceptNil()
	}
	check(receiver == nil && effects == 5)
	iface = nil
	effects = 0
	receiver = nil
	if iface != nil {
		receiver = iface.acceptNil()
	}
	check(receiver == nil && effects == 0)
	var outer **int
	v := 42
	if outer != nil && *outer != nil {
		v = **outer
	}
	check(v == 42)
	inner := (*int)(nil)
	outer = &inner
	v = 42
	if outer != nil && *outer != nil {
		v = **outer
	}
	check(v == 42)
	value := 0
	inner = &value
	v = 42
	if outer != nil && *outer != nil {
		v = **outer
	}
	check(v == 0)
	v = 42
	if absent != nil && absent.Nested != nil && *absent.Nested != nil {
		v = **absent.Nested
	}
	check(v == 42)
	h.Nested = outer
	v = 42
	if h != nil && h.Nested != nil && *h.Nested != nil {
		v = **h.Nested
	}
	check(v == 0)
	check(len(genericNil([]int(nil), []int{1})) == 1)
	check(len(genericNil(map[int]int(nil), map[int]int{1: 1})) == 1)
	present := &node{N: 7}
	check(asInterface((*node)(nil), present) == present)
	effects = 0
	r, e := propagateChain(nil, true)
	check(r == nil && e == nil && effects == 9)
	effects = 0
	r, e = propagateChain(present, true)
	check(r == nil && e != nil && effects == 69)
	effects = 0
	r, e = propagateChain(present, false)
	check(r == present && e == nil && effects == 69)
	effects = 0
	check(handlerChain(nil, true) == nil && effects == 0)
	effects = 0
	check(handlerChain(present, true) == nil && effects == 67)
	effects = 0
	check(statementChain(nil, true) == nil && effects == 0)
	effects = 0
	check(statementChain(present, true) != nil && effects == 6)
	var optional func(int) int
	if optional == nil {
		optional = func(x int) int { return x + 1 }
	}
	check(optional(2) == 3)
	effects = 0
	var funcNode *node
	if funcNode == nil {
		if funcNode == nil {
			funcNode = &node{N: effect(1)}
		}
	}
	check(funcNode.N == 1 && effects == 1)
	effects = 0
	var stored *node
	if stored == nil {
		stored = present
	}
	if stored == nil {
		stored = &node{N: effect(2)}
	}
	check(stored == present && effects == 0)
	type slots struct{ P *node }
	obj := &slots{P: present}
	if obj.P == nil {
		obj.P = &node{N: effect(3)}
	}
	check(obj.P == present && effects == 0)
}
