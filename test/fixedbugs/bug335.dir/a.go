package a

type T interface{}

func f() T { return nil }

var Foo T = f()
