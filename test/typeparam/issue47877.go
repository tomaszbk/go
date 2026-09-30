// run


package main

type Map[K comparable, V any] struct {
        m map[K]V
}

func NewMap[K comparable, V any]() Map[K, V] {
        return Map[K, V]{m: map[K]V{}}
}

func (m Map[K, V]) Get(key K) V {
        return m.m[key]
}

func main() {
        _ = NewMap[int, struct{}]()
}
