package a

import "cmd/compile/internal/loopvar/testdata/inlines/b"

func F() []*int {
	var s []*int
	for i := 0; i < 10; i++ {
		s = append(s, &i)
	}
	return s
}

func Fb() []*int {
	bf, _ := b.F()
	return bf
}
