// -lang=go1.20

package p

func F[P any, Q *P](p P) {}

var _ = F[int]

func G[R any](func(R)) {}

func _() {
	G(F[int])
}
