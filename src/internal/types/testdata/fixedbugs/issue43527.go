package p

const L = 10

type (
	_        [L]struct{}
	_        [A /* ERROR "undefined array length A or missing type constraint" */ ]struct{}
	_        [B /* ERROR "invalid array length B" */ ]struct{}
	_[A any] struct{}

	B int
)
