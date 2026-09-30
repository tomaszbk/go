// -lang=go1.26

package p

type Foo struct {
	Bar
}

type Bar struct {
	Baz int
}

var _ = Foo{Baz /* ERROR "use of promoted field Bar.Baz in struct literal of type Foo requires go1.27 or later" */ : 1}
