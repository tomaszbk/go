package a

import "unsafe"

type HookFunc func(x uint64)

var HookV unsafe.Pointer

func Hook(x uint64) {
	(*(*HookFunc)(HookV))(x)
}
