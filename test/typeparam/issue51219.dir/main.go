package main

import (
	"./a"
	"fmt"
)

func main() {
	var x a.I[a.JsonRaw]

	fmt.Printf("%v\n", x)
}
