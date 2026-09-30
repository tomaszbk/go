package p

import "C"

type T = T // ERROR HERE

//export F
func F(p *T) {}
