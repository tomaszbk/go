package a

import "sync"

type Val[T any] struct {
	mu  sync.RWMutex
	val T
}

func (v *Val[T]) Has() {
	v.mu.RLock()
}
