// build

package main

type T int

const K T = 5

type P struct {
	a [K]*byte
}

//go:noinline
func f(p *P) {
	for i := range K {
		p.a[i] = nil
	}
}
func main() {
	f(nil)
}
