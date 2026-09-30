// compile

// Ensure that late expansion correctly handles an OpIData with type interface{}

package p

type S struct{}

func (S) M() {}

type I interface {
	M()
}

func f(i I) {
	o := i.(interface{})
	if _, ok := i.(*S); ok {
		o = nil
	}
	println(o)
}
