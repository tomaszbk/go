package p

import "math/big"

// From go.dev/issue/18419
func _(x *big.Float) {
	x.form /* ERROR "x.form undefined (cannot refer to unexported field form)" */ ()
}

// From go.dev/issue/31053
func _() {
	_ = big.Float{form /* ERROR "cannot refer to unexported field form in struct literal of type big.Float" */ : 0}
}
