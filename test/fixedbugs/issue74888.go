// compile

package main

type P struct {
	q struct{}
	p *int
}

func f(x any) {
	h(x.(P))
}

//go:noinline
func h(P) {
}
