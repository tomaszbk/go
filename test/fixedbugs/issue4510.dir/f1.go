package p

import "fmt" // GCCGO_ERROR "fmt redeclared|imported"

var _ = fmt.Printf
