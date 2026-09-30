// run


package main

func main() {
	x := []uint{0}
	x[0] &^= f()
}

func f() uint {
	return 1<<31 // doesn't panic with 1<<31 - 1
}
