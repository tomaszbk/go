// run


// issue 5056: escape analysis not applied to wrapper functions

package main

type Foo int16

func (f Foo) Esc() *int{
	x := int(f)
	return &x
}

type iface interface {
	Esc() *int
}

var bar, foobar *int

func main() {
	var quux iface
	var x Foo
	
	quux = x
	bar = quux.Esc()
	foobar = quux.Esc()
	if bar == foobar {
		panic("bar == foobar")
	}
}
