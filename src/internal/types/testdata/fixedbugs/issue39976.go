package p

type policy[K, V any] interface{}
type LRU[K, V any] struct{}

func NewCache[K, V any](p policy[K, V]) {}

func _() {
	var lru LRU[int, string]
	NewCache[int, string](&lru)
	NewCache /* ERROR "cannot infer K" */ (&lru)
}
