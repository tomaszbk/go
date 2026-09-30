//go:build cgo && !netgo && ((linux && !android) || netbsd || solaris)

package net

/*
#include <sys/types.h>
#include <sys/socket.h>

#include <netdb.h>
*/
import "C"

import "unsafe"

func cgoNameinfoPTR(b []byte, sa *C.struct_sockaddr, salen C.socklen_t) (int, error) {
	gerrno, err := C.getnameinfo(sa, salen, (*C.char)(unsafe.Pointer(&b[0])), C.socklen_t(len(b)), nil, 0, C.NI_NAMEREQD)
	return int(gerrno), err
}
