// run


package main

const targetPC = uintptr(0xdeadbeef)

type payload struct {
	x uintptr
	y *uintptr
	fn   [2]func()
}

var p payload
var v []byte

func init() {
	p.x = targetPC
	p.y = &p.x
	p.fn[0] = func() {}
	p.fn[1] = func() {}
}

//go:noinline
func trigger(n int) {
	defer func() { recover() }()

	if n < len(p.fn) {
		p.fn[n&1]()

		s := make([]byte, n)
		v = s
	}
}

func main() {
	trigger(-1)
}
