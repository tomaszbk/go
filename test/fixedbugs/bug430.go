// compile

// gccgo crashed compiling this.

package main

type S struct {
	f [2][]int
}

func F() (r [2][]int) {
	return
}

func main() {
	var a []S
	a[0].f = F()
}
