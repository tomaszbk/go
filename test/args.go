// run arg1 arg2

// Test os.Args.

package main

import "os"

func main() {
	if len(os.Args) != 3 {
		panic("argc")
	}
	if os.Args[1] != "arg1" {
		panic("arg1")
	}
	if os.Args[2] != "arg2" {
		panic("arg2")
	}
}
