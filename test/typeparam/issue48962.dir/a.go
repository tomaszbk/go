package a

type (
	A[P any]               [10]P
	S[P any]               struct{ f P }
	P[P any]               *P
	M[K comparable, V any] map[K]V
)
