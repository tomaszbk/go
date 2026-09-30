//go:build cgo && !netgo

package net

//#include <netdb.h>
import "C"

const cgoAddrInfoFlags = C.AI_CANONNAME
