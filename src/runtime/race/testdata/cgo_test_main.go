package main

/*
int sync;

void Notify(void)
{
	__sync_fetch_and_add(&sync, 1);
}

void Wait(void)
{
	while(__sync_fetch_and_add(&sync, 0) == 0) {}
}
*/
import "C"

func main() {
	data := 0
	go func() {
		data = 1
		C.Notify()
	}()
	C.Wait()
	_ = data
}
