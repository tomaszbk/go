// asmcheck

package codegen

func main() {
	// amd64:"LEAQ command-line-arguments\\.h\\.func1"
	f(1)()

	// amd64:"LEAQ command-line-arguments\\.h\\.func1"
	g(2)()

	// amd64:"LEAQ command-line-arguments\\.h\\.func1"
	h(3)()

	// amd64:"LEAQ command-line-arguments\\.h\\.func1"
	f(4)()
}

func f(x int) func() {
	// amd64:"LEAQ command-line-arguments\\.h\\.func1"
	return g(x)
}

func g(x int) func() {
	// amd64:"LEAQ command-line-arguments\\.h\\.func1"
	return h(x)
}

func h(x int) func() {
	// amd64:"LEAQ command-line-arguments\\.h\\.func1"
	return func() { defer func() {}() } // defer prevents inlining
}
