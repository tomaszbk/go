package cgotest

// #include <stddef.h>
import "C"

func generic[T, U any](t T, u U) {}

func useGeneric() {
	const zero C.size_t = 0

	generic(zero, zero)
	generic[C.size_t, C.size_t](0, 0)
}
