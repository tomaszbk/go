// compile

package p

func f() {
	var st struct {
		s string
		i int16
	}
	_ = func() {
		var m map[int16]int
		m[st.i] = 0
	}
}
