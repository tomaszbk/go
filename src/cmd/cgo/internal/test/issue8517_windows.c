#include "windows.h"

extern void testHandleLeaksCallback();

DWORD WINAPI testHandleLeaksFunc(LPVOID lpThreadParameter)
{
	int i;
	for(i = 0; i < 100; i++) {
		testHandleLeaksCallback();
	}
	return 0;
}

void testHandleLeaks()
{
	HANDLE h;
	h = CreateThread(NULL, 0, &testHandleLeaksFunc, 0, 0, NULL);
	WaitForSingleObject(h, INFINITE);
	CloseHandle(h);
}
