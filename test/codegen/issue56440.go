// asmcheck

// Check to make sure that we recognize when the length of an append
// is constant. We check this by making sure that the constant length
// is folded into a load offset.

package codegen

func f(x []int) int {
	s := make([]int, 3)
	s = append(s, 4, 5)
	// amd64:`MOVQ 40\(.*\),`
	return x[len(s)]
}

func g(x []int, p *bool) int {
	s := make([]int, 3)
	for {
		s = s[:3]
		if cap(s) < 5 {
			s = make([]int, 3, 5)
		}
		s = append(s, 4, 5)
		if *p {
			// amd64:`MOVQ 40\(.*\),`
			return x[len(s)]
		}
	}
	return 0
}
