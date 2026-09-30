// compile


package main

type A[P any] = struct{ _ P }

type N[P any] A[P]

func f[P any](N[P]) {}

var _ = f[int]
