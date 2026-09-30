package main

// extern int CFunc(void);
import "C"

import (
	"fmt"
	"os"
)

func main() {
	got := C.CFunc()
	const want = (1 << 8) | 2
	if got != want {
		fmt.Printf("got %#x, want %#x\n", got, want)
		os.Exit(1)
	}
}
