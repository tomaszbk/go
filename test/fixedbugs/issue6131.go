// compile


// Issue 6131: missing typecheck after reducing
// n%1 == 0 to a constant value.

package main

func isGood(n int) bool {
	return n%1 == 0
}

func main() {
	if !isGood(256) {
		panic("!isGood")
	}
}
