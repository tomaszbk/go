package p

type A = [4]int
type B = map[string]interface{}

func _[T ~A](x T) {
	_ = len(x)
}

func _[U ~A](x U) {
	_ = cap(x)
}

func _[V ~A]() {
	_ = V{}
}

func _[W ~B](a interface{}) {
	_ = a.(W)["key"]
}
