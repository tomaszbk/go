// run


package main

import (
	"fmt"
	"os"
)

var (
	e interface{}
	s = struct{ a *int }{}
	b = e == s
)

func test(obj interface{}) {
	if obj != struct{ a *string }{} {
	}
}

var x int

func f() [2]string {
	x++
	return [2]string{"abc", "def"}
}

func main() {
	var e interface{} = [2]string{"abc", "def"}
	_ = e == f()
	if x != 1 {
		fmt.Println("x=", x)
		os.Exit(1)
	}
}
