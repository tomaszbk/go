package foo_test

import (
	"fmt"

	"example.com/error"
)

func Print(s string) {
	fmt.Println(s)
}

func Example() {
	Print(error.Hello)
}
