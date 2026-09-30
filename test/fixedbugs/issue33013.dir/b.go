package b

import "./a"

type Service uint64
type ServiceDesc struct {
	X int
	uc
}

type uc interface {
	f() a.G
}

var q int

func RS(svcd *ServiceDesc, server interface{}, qq uint8) *Service {
	defer func() { q += int(qq) }()
	return nil
}
