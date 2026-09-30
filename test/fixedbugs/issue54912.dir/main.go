package main

import "test/a"

func main() {
	a.F(new(int), 0)()
}
