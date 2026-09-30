#include <pthread.h>

/*
 * Call pthread_create, retrying on EAGAIN.
 */
extern int _cgo_try_pthread_create(pthread_t*, const pthread_attr_t*, void* (*)(void*), void*);

extern void* threadentry(void*);
