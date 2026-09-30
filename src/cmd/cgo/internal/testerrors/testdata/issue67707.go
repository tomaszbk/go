package p

import "C"

func F() *C.char {
	s, err := C.CString("hi") // ERROR HERE: no two-result form
	if err != nil {
		println(err)
	}
	return s
}
