package b

import "./a"

type (
	ABuilder = a.Builder
)

func Bfunc() ABuilder {
	return ABuilder{}
}
