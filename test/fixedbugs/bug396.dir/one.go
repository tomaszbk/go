package one

// Issue 2687
type T struct { int }

func New(i int) T { return T{i} }
