// errorcheck


// issue 1556
package foo

type I interface {
	m() int
}

type T int

var _ I = T(0)	// GCCGO_ERROR "incompatible"

func (T) m(buf []byte) (a int, b xxxx) {  // ERROR "xxxx"
	return 0, nil
}
