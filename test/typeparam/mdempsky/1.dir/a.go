package a

type T[_ any] int

func F() { _ = new(T[int]) }
