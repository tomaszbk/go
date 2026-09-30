//go:build unix

#include "libcgo.h"

#include <stdlib.h>

/* Stub for calling setenv */
void
x_cgo_setenv(char **arg)
{
	_cgo_tsan_acquire();
	setenv(arg[0], arg[1], 1);
	_cgo_tsan_release();
}

/* Stub for calling unsetenv */
void
x_cgo_unsetenv(char **arg)
{
	_cgo_tsan_acquire();
	unsetenv(arg[0]);
	_cgo_tsan_release();
}
