// compile

package p

func f_ssa(x int, p *int) {
	if false {
		y := x + 5
		for {
			*p = y
		}
	}
}
