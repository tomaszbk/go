package main

import (
	bad "issue46653.dir/bad"
)

func main() {
	bad.Bad()
}

func neverCalled() L {
	m := make(map[string]L)
	return m[""]
}

type L struct {
	A Data
	B Data
}

type Data struct {
	F1 [22][]string
}
