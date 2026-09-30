// compile

// Ensure that late expansion correctly handles an OpLoad with type interface{}

package p

type iface interface {
	m()
}

type it interface{}

type makeIface func() iface

func f() {
	var im makeIface
	e := im().(it)
	_ = &e
}
