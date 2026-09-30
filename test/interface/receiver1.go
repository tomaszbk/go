// errorcheck

// Verify compiler complains about missing implicit methods.
// Does not compile.

package main

type T int

func (t T) V()
func (t *T) P()

type V interface {
	V()
}
type P interface {
	P()
	V()
}

type S struct {
	T
}
type SP struct {
	*T
}

func main() {
	var t T
	var v V
	var p P
	var s S
	var sp SP

	v = t
	p = t // ERROR "does not implement|requires a pointer|cannot use"
	_, _ = v, p
	v = &t
	p = &t
	_, _ = v, p

	v = s
	p = s // ERROR "does not implement|requires a pointer|cannot use"
	_, _ = v, p
	v = &s
	p = &s
	_, _ = v, p

	v = sp
	p = sp // no error!
	_, _ = v, p
	v = &sp
	p = &sp
	_, _ = v, p
}
