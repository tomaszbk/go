// run

package main

import "strings"

//go:noinline
func f(i uint8) byte {
	return ""[i>>9]
}

func main() {
	defer func() {
		r := recover()
		if r == nil {
			panic("missing bounds panic")
		}
		if got := r.(error).Error(); !strings.Contains(got, "index out of range") {
			panic(got)
		}
	}()
	_ = f(7)
}
