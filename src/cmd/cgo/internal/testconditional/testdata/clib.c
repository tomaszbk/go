#include "clib.h"
#include <errno.h>
static int calls;
int c_step(int n) { calls++; return n; }
int c_sum(c_buf *b, int n) { calls++; return (b ? b->n : 0) + n; }
int c_div(int a, int b) { calls++; if (!b) { errno = EDOM; return -1; } return a / b; }
int c_calls(void) { return calls; }
void c_reset(void) { calls = 0; }
