// -lang=go1.13


package p

import "io"

// A "duplicate method" error should be reported for
// all these interfaces, irrespective of which package
// the embedded Reader is coming from.

type _ interface {
	Reader
	Reader // ERROR "duplicate method Read"
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type _ interface {
	io.Reader
	Reader // ERROR "duplicate method Read"
}

type _ interface {
	io.Reader
	io /* ERROR "duplicate method Read" */ .Reader
}
