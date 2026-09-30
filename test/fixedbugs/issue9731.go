// compile


package p

func f(x interface{}) {
	switch x := x.(type) {
	case int:
		func() {
			_ = x
		}()
	case map[int]int:
		func() {
			for range x {
			}
		}()
	}
}
