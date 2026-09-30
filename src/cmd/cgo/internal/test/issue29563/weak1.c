extern int weaksym __attribute__((__weak__));
int weaksym = 42;

int foo1()
{
	return weaksym;
}
