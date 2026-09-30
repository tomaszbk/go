// errorcheck


package p

type S struct{}

func (S) _[_ any]() {}

type _ interface {
	m[_ any]() // ERROR "interface method must have no type parameters"
}
