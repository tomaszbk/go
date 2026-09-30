// asmcheck


package codegen

// Make sure we use ADDQ instead of LEAQ when we can.

func f(p *[4][2]int, x int) *int {
	// amd64:"ADDQ" -"LEAQ"
	return &p[x][0]
}
