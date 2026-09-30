//go:build cgo && !netgo

package net

/*
#cgo LDFLAGS: -lsocket -lnsl
#include <netdb.h>
*/
import "C"

const cgoAddrInfoFlags = C.AI_CANONNAME | C.AI_V4MAPPED | C.AI_ALL
