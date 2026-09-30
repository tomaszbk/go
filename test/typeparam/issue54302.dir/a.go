package a

func A() {
	B[int](new(G[int]))
}

func B[T any](iface interface{ M(T) }) {
	x, ok := iface.(*G[T])
	if !ok || iface != x {
		panic("FAIL")
	}
}

type G[T any] struct{}

func (*G[T]) M(T) {}
