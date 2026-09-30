package a

type S[T any] struct {
	F T
}

var X = S[int]{}
