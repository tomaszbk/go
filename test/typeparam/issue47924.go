// compile


package p

type Cache[K any] struct{}

func (c Cache[K]) foo(x interface{}, f func(K) bool) {
	f(x.(K))
}

var _ Cache[int]
