package main

import "./a"

type S struct{}

func (*S) F() *S { return nil }

func main() {
	var _ a.I[*S] = &S{}
}
