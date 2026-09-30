package a

type Option[T any] interface {
	ToSeq() Seq[T]
}

type Seq[T any] []T

func (r Seq[T]) Find(p func(v T) bool) Option[T] {
	panic("")
}
