// compile -d=ssa/check/on

package p

func f()

func g() {
	var a []int
	var b bool
	for {
		b = (b && b) != (b && b)
		for b && b == b || true {
			f()
			_ = a[0]
		}
		_ = &b
		a = []int{}
	}
}
