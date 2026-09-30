// compile

package p

type Foo[T any] struct {
}

func (foo Foo[T]) Get()  {
}

var(
	_ = Foo[byte]{}
	_ = Foo[[]byte]{}
	_ = Foo[map[byte]rune]{}

	_ = Foo[rune]{}
	_ = Foo[[]rune]{}
	_ = Foo[map[rune]byte]{}
)
