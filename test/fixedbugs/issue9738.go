// run

package main

func F() (x int) {
	defer func() {
		if x != 42 {
			println("BUG: x =", x)
		}
	}()
	return 42
}

func main() {
	F()
}
