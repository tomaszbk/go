// errorcheck


package main

type T struct{}
type P *T

func (t *T) Meth() {}
func (t T) Meth2() {}

func main() {
	t := &T{}
	p := P(t)
	p.Meth()  // ERROR "undefined"
	p.Meth2() // ERROR "undefined"
}
