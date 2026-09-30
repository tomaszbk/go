package c

import (
	"./a"
	"./b"
)

type BI interface {
	Something(s int64) int64
	Another(pxp a.G) int32
}

func BRS(sd *b.ServiceDesc, server BI, xyz int) *b.Service {
	return b.RS(sd, server, 7)
}
