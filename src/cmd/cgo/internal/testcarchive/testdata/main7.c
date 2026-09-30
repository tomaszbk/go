// Test that lots of calls don't deadlock.

#include <stdio.h>

#include "libgo7.h"

int main() {
	int i;

	for (i = 0; i < 100000; i++) {
		GoFunction7();
	}
	return 0;
}
