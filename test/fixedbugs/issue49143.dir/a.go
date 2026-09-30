package a

import "sync"

type Loader[K comparable, R any] struct {
	batch *LoaderBatch[K, R]
}

func (l *Loader[K, R]) Load() error {
	l.batch.f()
	return nil
}

type LoaderBatch[K comparable, R any] struct {
	once    *sync.Once
}

func (b *LoaderBatch[K, R]) f() {
	b.once.Do(func() {})
}
