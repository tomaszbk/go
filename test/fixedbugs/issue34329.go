// errorcheck -lang=go1.13

package p

type I interface{ M() }

type _ interface {
	I
	I // ERROR "duplicate method M"
}
