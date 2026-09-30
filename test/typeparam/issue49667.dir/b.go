package b

import "./a"

type B[T any] struct {
	_ a.A[T]
}
