//go:build go1.20
// +build go1.20

package a

var (
	// Okay directive wise but the compiler will complain that
	// imports must appear before other declarations.
	//go:embed embedText // ok
	foo string
)

import (
	"fmt"

	_ "embed"
)

// This is main function
func main() {
	fmt.Println(s)
}
