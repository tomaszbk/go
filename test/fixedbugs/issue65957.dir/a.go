package a

var s any

//go:noinline
func F() {
	s = new([4]int32)
}
