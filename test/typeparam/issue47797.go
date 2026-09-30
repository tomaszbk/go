// compile


package p

type Foo[T any] struct {
        Val T
}

func (f Foo[T]) Bat() {}

type Bar struct {
        Foo[int]
}

func foo() {
        var b Bar
        b.Bat()
}
