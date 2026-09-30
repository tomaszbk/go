// compile


// Ensure that late expansion correctly set OpLoad argument type interface{}

package p

type iface interface {
	m()
}

type it interface{}

type makeIface func() iface

func f() {
	var im makeIface
	e := im().(it)
	g(e)
}

//go:noinline
func g(i it) {}
