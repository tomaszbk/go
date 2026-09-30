// run


// issue 8039. defer copy(x, <-c) did not rewrite <-c properly.

package main

func f(s []int) {
	c := make(chan []int, 1)
	c <- []int{1}
	defer copy(s, <-c)
}

func main() {
	x := make([]int, 1)
	f(x)
	if x[0] != 1 {
		println("BUG", x[0])
	}
}
