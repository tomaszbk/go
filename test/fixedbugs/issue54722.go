// compile

package p

type G[T any] struct {
	h H[G[T]]
}

type H[T any] struct{}

var x G[int]
