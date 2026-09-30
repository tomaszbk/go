// run

package main

var a [10]int
var b [1e1]int

func main() {
	if len(a) != 10 || len(b) != 10 {
		println("len", len(a), len(b))
		panic("fail")
	}
}
