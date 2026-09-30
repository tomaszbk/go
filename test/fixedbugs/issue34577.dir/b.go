package b

import "./a"

type B struct {
	s string
}

func (b B) Func(x a.A) a.A {
	return a.W(x, k, b)
}

type ktype int

const k ktype = 0

func Func2() a.AI {
	return a.ACC
}
