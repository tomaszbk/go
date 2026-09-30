package main

import "./a"

type No struct {
	a.EDO
}

func X() No {
	return No{}
}

func main() {
	X()
}
