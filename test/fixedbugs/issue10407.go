// runoutput


// Issue 10407: gccgo failed to remove carriage returns
// from raw string literals.

package main

import "fmt"

func main() {
	fmt.Println("package main\nfunc main() { if `a\rb\r\nc` != \"ab\\nc\" { panic(42) }}")
}
