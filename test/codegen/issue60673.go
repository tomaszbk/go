// asmcheck

package codegen

//go:noinline
func f(x int32) {
}

func g(p *int32) {
	// argument marshaling code should live at line 17, not line 15.
	x := *p
	// 386: `MOVL [A-Z]+, \(SP\)`
	f(x)
}
