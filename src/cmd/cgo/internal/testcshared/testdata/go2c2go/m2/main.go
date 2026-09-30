package main

// #include "libtestgo2c2go.h"
import "C"

import (
	"fmt"
	"os"
)

func main() {
	got := C.GoFunc()
	const want = 1
	if got != want {
		fmt.Printf("got %#x, want %#x\n", got, want)
		os.Exit(1)
	}
}
