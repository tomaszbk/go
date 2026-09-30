package main

import "./a"

type S[Idx any] struct {
	A string
	B Idx
}

type O[Idx any] struct {
	A int
	B a.I[Idx]
}
