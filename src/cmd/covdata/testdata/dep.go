package dep

func Dep1() int {
	return 42
}

func PDep(x int) {
	if x != 1010101 {
		println(x)
	} else {
		panic("bad")
	}
}
