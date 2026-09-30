// run


package main

var c chan int

func main() {
	c = make(chan int);
	go func() { c <- 0 } ();
	<-c
}
