package main

import (
        "reflect"
        fake "./a" // 2nd package with name "reflect"
)

type T struct {
        _ fake.Type
}

func (T) f()            {}
func (T) G() (_ int)    { return }
func (T) H() (_, _ int) { return }

func main() {
        var x T
        typ := reflect.TypeOf(x)
        for i := 0; i < typ.NumMethod(); i++ {
                _ = typ.Method(i) // must not crash
        }
}
