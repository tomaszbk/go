// build


package main

const a = 0

func f() {
	const a = 5
}

func main() {
	if a != 0 {
		println("a=", a)
		panic("fail")
	}
}
