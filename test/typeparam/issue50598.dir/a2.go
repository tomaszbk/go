package a2

import "./a0"

func New() int {
	return a0.Builder[int]{}.New1()
}
