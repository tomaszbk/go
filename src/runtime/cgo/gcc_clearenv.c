//go:build linux

#include "libcgo.h"

#include <stdlib.h>

/* Stub for calling clearenv */
void
x_cgo_clearenv(void **env __attribute__((unused)))
{
	_cgo_tsan_acquire();
	clearenv();
	_cgo_tsan_release();
}
