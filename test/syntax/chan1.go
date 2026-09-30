// errorcheck


package main

var c chan int
var v int

func main() {
	if c <- v { // ERROR "cannot use c <- v as value|send statement used as value"
	}
}

var _ = c <- v // ERROR "unexpected <-|send statement used as value"
