// run

package main

func f[B any](b B) {
	if b1, ok := any(b).(interface{ m1() }); ok {
		panic(1)
		_ = b1.(B)
	}
	if b2, ok := any(b).(interface{ m2() }); ok {
		panic(2)
		_ = b2.(B)
	}
}

type S struct{}

func (S) m3() {}

func main() {
	f(S{})
}
