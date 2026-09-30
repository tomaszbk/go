package cache

// assert panics with the given msg if cond is not true.
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
