#include "_cgo_export.h"

static void callDestructorCallback() {
	GoDestructorCallback();
}

static void (*destructorFn)(void);

void registerDestructor() {
	destructorFn = callDestructorCallback;
}

__attribute__((destructor))
static void destructor() {
	if (destructorFn) {
		destructorFn();
	}
}
