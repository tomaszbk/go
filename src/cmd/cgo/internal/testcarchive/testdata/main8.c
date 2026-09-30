// Test preemption.

#include <stdlib.h>

#include "libgo8.h"

int main() {
	GoFunction8();

	// That should have exited the program.
	abort();
}
