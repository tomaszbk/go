// run

// issue 1172

package main

func main() {
	var i interface{}
	c := make(chan int, 1)
	c <- 1
	select {
	case i = <-c: // error on this line
	}
	if i != 1 {
		println("bad i", i)
		panic("BUG")
	}
}
