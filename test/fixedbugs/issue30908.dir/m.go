package main

import (
	"os"

	"./b"
)

func main() {
	seed := "some things are better"
	bsl := []byte(seed)
	b.CallReadValues("/dev/null")
	vals, err := b.ReadValues(bsl)
	if vals["better"] != seed || err != nil {
		os.Exit(1)
	}
}
