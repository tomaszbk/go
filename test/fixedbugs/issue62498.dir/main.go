package main

import "./a"

func main() {
	a.One(nil)
	Two(nil)
}

func Two(L any) {
	func() {
		defer a.F(L)
	}()
}
