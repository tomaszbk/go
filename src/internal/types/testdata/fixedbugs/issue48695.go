package p

func g[P ~func(T) P, T any](P) {}

func _() {
	type F func(int) F
	var f F
	g(f)
	_ = g[F]
}
