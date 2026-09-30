// run


// Issue 8961. Empty composite literals to small globals were not filled in
package main

type small struct { a int }
var foo small

func main() {
	foo.a = 1
	foo = small{}
	if foo.a != 0 {
		println("expected foo.a to be 0, was", foo.a)
		panic("composite literal not filled in")
	}
}
