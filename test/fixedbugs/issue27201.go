// run


package main

import (
	"runtime"
	"strings"
)

func main() {
	f(nil)
}

func f(p *int32) {
	defer checkstack()
	v := *p         // panic should happen here, line 17
	sink = int64(v) // not here, line 18
}

var sink int64

func checkstack() {
	_ = recover()
	var buf [1024]byte
	n := runtime.Stack(buf[:], false)
	s := string(buf[:n])
	if strings.Contains(s, "issue27201.go:18 ") {
		panic("panic at wrong location")
	}
	if !strings.Contains(s, "issue27201.go:17 ") {
		panic("no panic at correct location")
	}
}
