// asmcheck


package codegen

func f(x, y int, p *int) {
	// amd64:`MOVQ AX, BX`
	h(8, x)
	*p = y
}

//go:noinline
func h(a, b int) {
}
