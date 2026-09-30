// errorcheck -lang=go1.17


package p

type empty interface{}

type Foo[T empty] int // ERROR "type parameter requires go1\.18 or later \(-lang was set to go1\.17; check go.mod\)"

func Bar[T empty]() {} // ERROR "type parameter requires go1\.18 or later \(-lang was set to go1\.17; check go.mod\)"
