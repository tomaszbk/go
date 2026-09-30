// run


package main

var n int

func f() int {
	n++
	return n
}

func main() {
	x := []int{0,1,2,3,4,5,6,7,8,9,10}
	n = 5
	y := x[f():f()]
	if len(y) != 1 || y[0] != 6 {
		println("BUG bug261", len(y), y[0])
	}
}
