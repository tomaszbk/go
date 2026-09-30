#include <errno.h>
#include <string.h>

#include "clib.h"

static int ncalls;
static int nhandled;

int c_div(int num, int den) {
	ncalls++;
	if (den == 0) {
		errno = EDOM;
		return -1;
	}
	return num / den;
}

void c_touch(int x) {
	ncalls++;
	if (x < 0) {
		errno = EINVAL;
	}
}

int c_add3(int a, int b, int c) {
	ncalls++;
	return a + b + c;
}

int c_strlen(const char *s) {
	ncalls++;
	if (*s == 0) {
		errno = EINVAL;
		return -1;
	}
	return (int)strlen(s);
}

int c_bufsum(c_buf *b, int k) {
	ncalls++;
	if (k == 0) {
		errno = EDOM;
		return -1;
	}
	return b->n + k;
}

void c_note_handled(void) { nhandled++; }
int c_nhandled(void) { return nhandled; }
int c_ncalls(void) { return ncalls; }

void c_reset(void) {
	ncalls = 0;
	nhandled = 0;
}
