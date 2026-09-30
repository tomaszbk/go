// run


// Issue 23545: gccgo didn't lower array comparison to
// proper equality function in some case.

package main

func main() {
	if a := Get(); a != dummyID(1234) {
		panic("FAIL")
	}
}

func dummyID(x int) [Size]interface{} {
	var out [Size]interface{}
	out[0] = x
	return out
}

const Size = 32

type OutputID [Size]interface{}

//go:noinline
func Get() OutputID {
	return dummyID(1234)
}
