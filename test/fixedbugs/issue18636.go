// run


package main

import "runtime/debug"

type Foo struct {
	A [1 << 20]byte
	B string
}

func run(c chan bool) {
	f := new(Foo)
	*f = Foo{B: "hello"}
	c <- true
}

func main() {
	debug.SetMaxStack(1 << 16)
	c := make(chan bool)
	go run(c)
	<-c
}
