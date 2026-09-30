package main

import (
	bad "issue47185.dir/bad"
)

func main() {
	another()
	bad.Bad()
}

func another() L {
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
