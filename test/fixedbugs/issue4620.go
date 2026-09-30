// run


// Issue 4620: map indexes are not evaluated before assignment of other elements

package main

import "fmt"

func main() {
	m := map[int]int{0:1}
	i := 0
	i, m[i] = 1, 2
	if m[0] != 2 {
		fmt.Println(m)
		panic("m[i] != 2")
	}
}
