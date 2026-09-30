package b

import (
	"./a"
)

type S struct{}

func (s *S) M1() a.I {
	return a.NewWithF(s.M2)
}

func (s *S) M2() {}
