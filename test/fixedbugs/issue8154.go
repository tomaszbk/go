// compile


// Issue 8154: cmd/5g: ICE in walkexpr walk.c

package main

func main() {
	c := make(chan int)
	_ = [1][]func(){[]func(){func() { <-c }}}
}
