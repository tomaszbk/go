package c

import (
	"./b"
)

func BuildInt() int {
	return b.IntBuilder.New()
}
