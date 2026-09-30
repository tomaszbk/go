// compile


package main

type S struct {
	a [1 << 16]byte
}

func f1() {
	p := &S{}
	_ = p
}

type T [1 << 16]byte

func f2() {
	p := &T{}
	_ = p
}
