// run

package main

const c1 = iota
const c2 = iota

const c3 = 0 + iota<<8
const c4 = 1 + iota<<8

func main() {
	if c1 != 0 {
		panic(c1)
	}
	if c2 != 0 {
		panic(c2)
	}

	if c3 != 0 {
		panic(c3)
	}
	if c4 != 1 {
		panic(c4)
	}

	const c5 = iota
	const c6 = iota

	if c5 != 0 {
		panic(c5)
	}
	if c6 != 0 {
		panic(c6)
	}
}
