package c

import "./b"

func F() interface{} {
	go func(){}() // make it non-inlineable
	return b.F()
}

func P() interface{} {
	go func(){}() // make it non-inlineable
	return b.P()
}
