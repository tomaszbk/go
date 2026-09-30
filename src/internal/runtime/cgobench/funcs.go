//go:build cgo

package cgobench

/*
static void empty() {
}

void go_empty_callback();

static void callback() {
	go_empty_callback();
}

*/
import "C"

func EmptyC() {
	C.empty()
}

func CallbackC() {
	C.callback()
}

//export go_empty_callback
func go_empty_callback() {
}

//go:noinline
func Empty() {
}
