// run


package main

import (
	"errors"
	"fmt"
	"os"
)

func crashOnErr(err error) int {
	if err != nil {
		panic(err)
	}
	return 10
}

func main() {
	defer func() {
		if recover() == nil {
			fmt.Println("failed to have expected panic")
			os.Exit(1)
		}
	}()

	s := make([]int, crashOnErr(errors.New("test error")))
	println("unreachable: len(s) =", len(s))
}
