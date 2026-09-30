package p

type C interface {
	comparable
	[2]any | int
}

func f[T C]() {}

func _() {
	_ = f[[ /* ERROR "[2]any does not satisfy C (C mentions [2]any, but [2]any is not in the type set of C)" */ 2]any]
}
