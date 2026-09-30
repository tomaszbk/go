package a

type T[U interface{ M() T[U] }] int
