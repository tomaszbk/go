// -lang=go1.16


//go:build go1.21

package main

import "slices"

func main() {
	_ = slices.Clone([]string{}) // no error should be reported here
}
