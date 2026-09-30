package p

type T[P any] struct{}

type _ interface{ int | T /* ERROR "cannot use generic type" */ }