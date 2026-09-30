// +build !plan9,!windows

#include <stdint.h>

uint32_t threadExited;

void setExited(void *x) {
	__sync_fetch_and_add(&threadExited, 1);
}
