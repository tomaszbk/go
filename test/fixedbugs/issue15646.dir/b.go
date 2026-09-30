package main

import "./a" // import must succeed

func main() {
	if a.F()(a.T{}) != "m" {
		panic(0)
	}
	if a.Fp()(nil) != "mp" {
		panic(1)
	}
}
