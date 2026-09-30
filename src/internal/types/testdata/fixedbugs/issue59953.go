package p

func _()         { f(g) }
func f[P any](P) {}
func g[Q int](Q) {}
