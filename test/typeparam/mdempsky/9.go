// compile


package a

func f[V any]() []V { return []V{0: *new(V)} }

func g() { f[int]() }
