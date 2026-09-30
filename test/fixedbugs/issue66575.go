// run

package main

var (
	v0 = initv0()
	v1 = initv1()
)

const c = "c"

func initv0() string {
	println("initv0")
	if c != "" { // have a dependency on c
		return ""
	}
	return ""
}

func initv1() string {
	println("initv1")
	return ""
}

func main() {
	// do nothing
}
