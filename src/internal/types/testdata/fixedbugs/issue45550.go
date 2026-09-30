package p

type Builder[T ~struct{ Builder[T] }] struct{}
type myBuilder struct {
	Builder[myBuilder]
}
