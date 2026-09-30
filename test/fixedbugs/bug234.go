// run

package main

func main() {
	c := make(chan int, 1)
	c <- 100
	x, ok := <-c
	if x != 100 || !ok {
		println("x=", x, " ok=", ok, " want 100, true")
		panic("fail")
	}
	close(c)
	x, ok = <-c
	if x != 0 || ok {
		println("x=", x, " ok=", ok, " want 0, false")
		panic("fail")
	}
}
