//go:build cgo && !netgo && (dragonfly || freebsd)

package net

/*
#include <netdb.h>
*/
import "C"

const cgoAddrInfoFlags = (C.AI_CANONNAME | C.AI_V4MAPPED | C.AI_ALL) & C.AI_MASK
