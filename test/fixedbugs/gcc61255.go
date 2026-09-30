// compile


// PR61255: gccgo failed to compile IncDec statements on variadic functions.

package main

func main() {
	append([]byte{}, 0)[0]++
}
