package p

func f[P interface{ m(R) }, R any]() {}

type T = interface { m(int) }

func _() {
	_ = f[T]
}
