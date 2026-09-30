// run


// Issue 2423

package main

func main() {
	var x interface{} = "hello"

	switch x {
	case "hello":
	default:
		println("FAIL")
	}
}
