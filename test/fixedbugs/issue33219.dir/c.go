package c

import (
	"./a"
	"./b"
)

type BI interface {
	Another(pxp a.A) int32
}

//go:noinline
func BRS(sd a.A, xyz int) *b.Service {
	x := b.Yes(sd, nil)
	return b.No(x, 1)
}
