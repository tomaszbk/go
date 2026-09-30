package p

func _() {
	NewS /* ERROR "cannot infer T" */ ().M()
}

type S struct {}

func NewS[T any]() *S { panic(0) }

func (_ *S /* ERROR "S is not a generic type" */ [T]) M()
