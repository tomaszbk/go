// run

// Torture test for goroutines.
// Make a lot of goroutines, threaded together, and tear them down cleanly.

package main

import (
	"os"
	"strconv"
)

func f(left, right chan int) {
	left <- <-right
}

func main() {
	var n = 10000
	if len(os.Args) > 1 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			print("bad arg\n")
			os.Exit(1)
		}
	}
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	<-leftmost
}
