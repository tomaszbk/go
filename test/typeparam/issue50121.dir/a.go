package a

import (
	"math/rand"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Builder[T Integer] struct{}

func (r Builder[T]) New() T {
	return T(rand.Int())
}

var IntBuilder = Builder[int]{}

func BuildInt() int {
	return IntBuilder.New()
}
