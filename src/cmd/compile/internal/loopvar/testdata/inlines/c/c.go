package c

//go:noinline
func F() []*int {
	var s []*int
	for i := 0; i < 10; i++ {
		s = append(s, &i)
	}
	return s
}
