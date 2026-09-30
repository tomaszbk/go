package p

import (
	"./a"
	"./b"
)

func f() {
	a.NewWithFuncI((&b.S{}).M1)
}
