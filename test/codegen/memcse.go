// asmcheck


// Test common subexpression elimination of loads around other operations.

package codegen

func loadsAroundMemEqual(p *int, s1, s2 string) (int, bool) {
	x := *p
	eq := s1 == s2
	y := *p
	// arm64:"MOVD ZR, R0"
	return x - y, eq
}

func loadsAroundVarDefAndZero(p *int) (int, *int) {
	x := *p
	var s [2]*int
	s[0] = p
	y := *p
	// arm64:"MOVD ZR, R0"
	return x - y, s[0]
}
