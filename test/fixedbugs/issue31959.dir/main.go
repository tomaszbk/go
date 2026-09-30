// run

// Check import package contains type alias in function
// with the same name with an export type not panic

package main

import (
	"fmt"

	"./a"
)

func main() {
	fmt.Println(a.T{})
	a.F()
}
