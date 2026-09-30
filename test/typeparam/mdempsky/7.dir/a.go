package a

type I[T any] interface{ M() T }

var X I[int]
