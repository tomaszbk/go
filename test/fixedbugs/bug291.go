// run


// https://golang.org/issue/915

package main

type T struct {
	x int
}

var t = &T{42}
var i interface{} = t
var tt, ok = i.(*T)

func main() {
	if tt == nil || tt.x != 42 {
		println("BUG")
	}
}
