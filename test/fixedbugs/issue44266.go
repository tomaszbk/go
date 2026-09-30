// errorcheck

package p

import "io"

type T1 interface {
	io.Reader
}

type T2 struct {
	io.SectionReader
}

type T3 struct { // ERROR "invalid recursive type: T3 refers to itself"
	T1
	T2
	parent T3
}
