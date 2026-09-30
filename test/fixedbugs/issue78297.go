// compile

package p

func G[U any]() (u U) { return }

//go:noinline
func H[U any]() (u U) { return }

func F[T ~*[1]byte]() {
	_ = G[T]()[:]
	_ = H[T]()[:]
}

var _ = F[*[1]byte]
