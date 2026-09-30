package p

type A = int

type T[P any] *A

var _ T[int]
