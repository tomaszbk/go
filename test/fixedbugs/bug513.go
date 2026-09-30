// run -race -gcflags=all=-d=checkptr=0

//go:build ((linux && amd64) || (linux && ppc64le) || (darwin && amd64) || (freebsd && amd64) || (netbsd && amd64) || (windows && amd64)) && cgo


// Although -race turns on -d=checkptr, the explicit -d=checkptr=0
// should override it.

package main

import "unsafe"

var v1 = new([2]int16)
var v2 *[3]int64

func main() {
	v2 = (*[3]int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*v1)[0]))))
}
