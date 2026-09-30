//go:build !android

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

static pthread_t thread;

static void* threadfunc(void* dummy) {
	while(1) {
		sleep(1);
	}
}

int StartThread() {
	return pthread_create(&thread, NULL, &threadfunc, NULL);
}

int CancelThread() {
	void *r;
	pthread_cancel(thread);
	pthread_join(thread, &r);
	return (r == PTHREAD_CANCELED);
}
