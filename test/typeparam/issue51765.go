// compile

package p

type empty[T any] struct{}

func (this *empty[T]) Next() (empty T, _ error) {
	return empty, nil
}

var _ = &empty[string]{}
