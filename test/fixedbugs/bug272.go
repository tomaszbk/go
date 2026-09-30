// run


// https://golang.org/issue/589

package main

func main() {	
	n := int64(100)
	x := make([]int, n)
	x[99] = 234;	
	z := x[n-1]
	if z != 234 {
		println("BUG")
	}
	n |= 1<<32
	defer func() {
		recover()
	}()
	z = x[n-1]
	println("BUG2")
}
