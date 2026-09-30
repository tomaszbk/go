package main

import (
	"fmt"
)

type Any struct {
	s string
	b int64
}

var Sg = []interface{}{
	Any{"a", 10},
}

func main() {
	fmt.Println(Sg[0])
}
