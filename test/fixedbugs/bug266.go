// run


package main

func f() int {
	defer func() {
		recover()
	}()
	panic("oops")
}

func g() int {	
	return 12345
}

func main() {
	g()	// leave 12345 on stack
	x := f()
	if x != 0 {
		panic(x)
	}
}
