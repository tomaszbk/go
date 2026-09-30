package p

type T[_ any] int

func f[_ any]() {}
func g[_, _ any]() {}

func _() {
	_ = f[T /* ERROR "without instantiation" */ ]
	_ = g[T /* ERROR "without instantiation" */ , T /* ERROR "without instantiation" */ ]
}
