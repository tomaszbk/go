#include "_cgo_export.h"

#if defined(WIN32) || defined(_AIX)
extern void setCallback(void *);
void init() {
	setCallback(goCallback);
}
#else
void init() {}
#endif
