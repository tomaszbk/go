package b

import "./a"

type Ap1[A, B any] struct {
	opt a.Option[A]
}

type Ap2[A, B any] struct {
	opt a.Option[A]
}
