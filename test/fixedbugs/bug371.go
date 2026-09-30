// errorcheck

// issue 2343

package main

type T struct{}

func (t *T) pm() {}
func (t T) m()   {}

func main() {
	p := &T{}
	p.pm()
	p.m()

	q := &p
	q.m()  // ERROR "requires explicit dereference|undefined"
	q.pm() // ERROR "requires explicit dereference|undefined"
}
