#include <process.h>
#include <stdlib.h>
#include <stdio.h>

void gopanic(void);

static unsigned int __attribute__((__stdcall__))
die(void* x)
{
	gopanic();
	return 0;
}

void
start(void)
{
	if(_beginthreadex(0, 0, die, 0, 0, 0) != 0)
		printf("_beginthreadex failed\n");
}
