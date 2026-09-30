// run


package main

func main() {
	_ = F[bool]

	var x string
	_ = G(x == "foo")
}

func F[T ~bool](x string) {
	var _ T = x == "foo"
}

func G[T any](t T) *T {
	return &t
}
