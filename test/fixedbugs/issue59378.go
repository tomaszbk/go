// compile


package p

func f() {
	F([]int{}, func(*int) bool { return true })
}

func F[S []E, E any](a S, fn func(*E) bool) {
	for _, v := range a {
		G(a, func(e E) bool { return fn(&v) })
	}
}

func G[E any](s []E, f func(E) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}
