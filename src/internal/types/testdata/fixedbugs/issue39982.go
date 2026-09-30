package p

type (
	T[_ any] struct{}
	S[_ any] struct {
		data T[*T[int]]
	}
)

func _() {
	_ = S[int]{
		data: T[*T[int]]{},
	}
}

// full test case from issue

type (
	Element[TElem any] struct{}

	entry[K comparable] struct{}

	Cache[K comparable] struct {
		data map[K]*Element[*entry[K]]
	}
)

func _() {
	_ = Cache[int]{
		data: make(map[int](*Element[*entry[int]])),
	}
}
