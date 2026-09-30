// errorcheck

package p

type I interface {
	F()
}

type T struct {
}

const _ = I((*T)(nil)) // ERROR "is not constant"

func (*T) F() {
}
