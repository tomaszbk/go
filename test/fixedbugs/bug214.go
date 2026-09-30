// compile


// Used to crash the compiler.
// https://golang.org/issue/88

package main

func main() {
	x := make(map[int]int, 10);
	x[0], x[1] = 2, 6;
}
