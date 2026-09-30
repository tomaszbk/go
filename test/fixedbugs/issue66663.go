// compile


package p

type Iterator[A any] func() (bool, A)

type Range[A any] interface {
	Blocks() Iterator[Block[A]]
}

type Block[A any] interface {
	Range[A]
}

type rangeImpl[A any] struct{}

func (r *rangeImpl[A]) Blocks() Iterator[Block[A]] {
	return func() (bool, Block[A]) {
		var a Block[A]
		return false, a
	}
}

func NewRange[A any]() Range[A] {
	return &rangeImpl[A]{}
}

type AddrImpl struct{}

var _ = NewRange[AddrImpl]()
