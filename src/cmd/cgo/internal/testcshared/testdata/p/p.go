package p

import "C"

//export FromPkg
func FromPkg() int32 { return 1024 }

//export Divu
func Divu(a, b uint32) uint32 { return a / b }
