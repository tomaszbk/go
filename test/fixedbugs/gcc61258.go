// run


// PR61258: gccgo crashed when deleting a zero-sized key from a map.

package main

func main() {
	delete(make(map[[0]bool]int), [0]bool{})
}
