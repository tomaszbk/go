// build

package main

func f1() {
	type T struct {
		x int
	}
}

func f2() {
	type T struct {
		x float64
	}
}

func main() {
	f1()
	f2()
}

/*
1606416576: conflicting definitions for main.T·bug167
bug167.6:	type main.T·bug167 struct { x int }
bug167.6:	type main.T·bug167 struct { x float64 }
*/
