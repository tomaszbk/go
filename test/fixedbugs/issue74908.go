// compile

package main

type Type struct {
	any
}

type typeObject struct {
	e struct{}
	b *byte
}

func f(b *byte) Type {
	return Type{
		typeObject{
			b: b,
		},
	}
}
