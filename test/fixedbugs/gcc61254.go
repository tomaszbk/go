// compile


// PR61254: gccgo failed to compile a slice expression with missing indices.

package main

func main() {
	[][]int{}[:][0][0]++
}
