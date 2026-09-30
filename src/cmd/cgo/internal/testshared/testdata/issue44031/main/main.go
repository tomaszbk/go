package main

import "testshared/issue44031/b"

type t int

func (t) m() {}

type i interface{ m() } // test that unexported method is correctly marked

var v interface{} = t(0)

func main() {
	b.F()
	v.(i).m()
}
