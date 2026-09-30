// compile


// issue 8073.
// was "internal compiler error: overflow: float64 integer constant"

package main

func main() {
	var x int
	_ = float64(x * 0)
}
