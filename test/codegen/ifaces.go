// asmcheck

package codegen

type I interface{ M() }

func NopConvertIface(x I) I {
	// amd64:-`.*runtime.convI2I`
	return I(x)
}

func NopConvertGeneric[T any](x T) T {
	// amd64:-`.*runtime.convI2I`
	return T(x)
}

var NopConvertGenericIface = NopConvertGeneric[I]

func ConvToM(x any) I {
	// amd64:`CALL runtime.typeAssert` `MOVL 16\(.*\)` `MOVQ 8\(.*\)(.*\*1)`
	// arm64:`CALL runtime.typeAssert` `LDAR` `MOVWU` `MOVD \(R.*\)\(R.*\)`
	return x.(I)
}

func e1(x any, y *int) bool {
	// amd64:-`.*faceeq` `SETEQ`
	// arm64:-`.*faceeq` `CSET EQ`
	return x == y
}

func e2(x any, y *int) bool {
	// amd64:-`.*faceeq` `SETEQ`
	// arm64:-`.*faceeq` `CSET EQ`
	return y == x
}

type E *int

func e3(x any, y E) bool {
	// amd64:-`.*faceeq` `SETEQ`
	// arm64:-`.*faceeq` `CSET EQ`
	return x == y
}

type T int

func (t *T) M() {}

func i1(x I, y *T) bool {
	// amd64:-`.*faceeq` `SETEQ`
	// arm64:-`.*faceeq` `CSET EQ`
	return x == y
}

func i2(x I, y *T) bool {
	// amd64:-`.*faceeq` `SETEQ`
	// arm64:-`.*faceeq` `CSET EQ`
	return y == x
}
