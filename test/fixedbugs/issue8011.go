// run

package main

func main() {
	c := make(chan chan int, 1)
	c1 := make(chan int, 1)
	c1 <- 42
	c <- c1
	x := <-<-c
	if x != 42 {
		println("BUG:", x, "!= 42")
	}
}
