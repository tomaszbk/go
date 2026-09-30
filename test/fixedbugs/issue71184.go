// compile


package x

func F[T int32]() {
	_ = G[*[0]T]()[:]
}

func G[T any]() (v T) {
	return
}

var _ = F[int32]
