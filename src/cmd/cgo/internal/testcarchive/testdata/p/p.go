package p

import "C"

//export FromPkg
func FromPkg() int32 { return 1024 }
