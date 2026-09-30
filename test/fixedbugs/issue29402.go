// run
  

// Issue 29402: wrong optimization of comparison of
// constant and shift on MIPS.

package main

//go:noinline
func F(s []int) bool {
	half := len(s) / 2
	return half >= 0
}

func main() {
	b := F([]int{1, 2, 3, 4})
	if !b {
		panic("FAIL")
	}
}
