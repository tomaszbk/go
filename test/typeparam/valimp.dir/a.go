package a

type Value[T any] struct {
	val T
}

// The noinline directive should survive across import, and prevent instantiations
// of these functions from being inlined.

//go:noinline
func Get[T any](v *Value[T]) T {
	return v.val
}

//go:noinline
func Set[T any](v *Value[T], val T) {
	v.val = val
}

//go:noinline
func (v *Value[T]) Set(val T) {
	v.val = val
}

//go:noinline
func (v *Value[T]) Get() T {
	return v.val
}
