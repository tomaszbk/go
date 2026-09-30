// run

package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stdout.Write(nil)
	if err != nil {
		fmt.Printf("BUG: os.Stdout.Write(nil) = %v\n", err)
	}
}
