// Issue 18452: show pos info in undefined name errors

package p

import (
	"C"
	"fmt"
)

func a() {
	fmt.Println("Hello, world!")
	C.function_that_does_not_exist() // ERROR HERE
	C.pi                             // ERROR HERE
}
