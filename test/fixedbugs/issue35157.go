// compile

package p

func f() {
	var i int
	var b *bool
	var s0, s1, s2 string

	if *b {
		s2 = s2[:1]
		i = 1
	}
	s1 = s1[i:-i+i] + s1[-i+i:i+2]
	s1 = s0[i:-i]
}
