//go:build gccgo

#include "_cgo_export.h"
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

/* Test calling panic from C.  This is what SWIG does.  */

extern void _cgo_panic(const char *);
extern void *_cgo_allocate(size_t);

void
callPanic(void)
{
	_cgo_panic("panic from C");
}
