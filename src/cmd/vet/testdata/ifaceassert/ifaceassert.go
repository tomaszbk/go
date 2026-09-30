package ifaceassert

type I1 interface{ M() int }
type I2 interface{ M() string }

func _(x I1) {
	_ = x.(I2) // ERROR "impossible type assertion"
}
