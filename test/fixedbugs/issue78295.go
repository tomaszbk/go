// run

package main

type B *struct{ A }
type A interface{ m(B) }

type s struct{}

func (s) m(b B) {}

func main() {
	var b B = new(struct{ A })
	b.A = s{}
	(*b).m(b)
}
