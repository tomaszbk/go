package gen

type Uint64 struct {
	valGP
}

var kindUint64 = &kind{typ: "Uint64", reg: regClassGP}

func ConstUint64(c uint64, name string) (y Uint64) {
	y.initOp(&op{op: "const", kind: y.kind(), c: c, name: name})
	return y
}

func (Uint64) kind() *kind {
	return kindUint64
}

func (Uint64) wrap(x *op) Uint64 {
	var y Uint64
	y.initOp(x)
	return y
}
