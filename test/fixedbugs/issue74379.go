// run


package main

import (
	"errors"
	"fmt"
	"os"
)

func crashOnErr(err error) bool {
	if err != nil {
		panic(err)
	}
	return false
}

func main() {
	defer func() {
		if recover() == nil {
			fmt.Println("failed to have expected panic")
			os.Exit(1)
		}
	}()
	fmt.Println(crashOnErr(errors.New("test error")))
}
