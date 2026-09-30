//go:build boringcrypto

package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Printf("SKIP with boringcrypto enabled\n")
	os.Exit(0)
}
