package main

import "./p"

var fn0 = (*p.T0).NilCheck
var fn1 = (*p.T1).NilCheck
var valueFn0 = (*p.ValueT0).ValueNilCheck
var valueFn1 = (*p.ValueT1).ValueNilCheck

func main() {
	test("offset zero", func() { fn0((*p.T0)(nil)) })
	test("offset non-zero", func() { fn1((*p.T1)(nil)) })
	test("value offset zero", func() { valueFn0((*p.ValueT0)(nil)) })
	test("value offset non-zero", func() { valueFn1((*p.ValueT1)(nil)) })
}

func test(name string, call func()) {
	p.Called = false
	defer func() {
		if recover() == nil {
			panic(name + ": want panic")
		}
		if p.Called {
			panic(name + ": called embedded method")
		}
	}()
	call()
}
