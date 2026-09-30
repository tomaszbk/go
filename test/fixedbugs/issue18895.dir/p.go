package p

func F() { // ERROR "can inline"
	var v t
	v.m() // ERROR "inlining call"
}

type t int

func (t) m() {} // ERROR "can inline"
