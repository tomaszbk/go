// -lang=go1.13


// Check Go language version-specific errors.

package p

// interface embedding

type I interface { m() }

type _ interface {
	m()
	I // ERROR "duplicate method m"
}

type _ interface {
	I
	I // ERROR "duplicate method m"
}
