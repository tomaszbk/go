package main

import "./a"

var B = [2]string{"world", "hello"}

func main() {
	if a.A[0] != B[1] {
		panic("bad hello")
	}
}
