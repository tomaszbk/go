package p

type T0[P any] struct {
	f P
}

type T1 /* ERROR "invalid recursive type" */ struct {
	_ T0[T1]
}
