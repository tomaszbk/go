// run -gcflags="-d=checkptr"

package main

import (
	"strings"
	"unsafe"
)

func main() {
	defer func() {
		err := recover()
		if err == nil {
			panic("expected panic")
		}
		if got := err.(error).Error(); !strings.Contains(got, "slice bounds out of range") {
			panic("expected panic slice out of bound, got " + got)
		}
	}()
	s := make([]int64, 100)
	p := unsafe.Pointer(&s[0])
	n := 1000

	_ = (*[10]int64)(p)[:n:n]
}
