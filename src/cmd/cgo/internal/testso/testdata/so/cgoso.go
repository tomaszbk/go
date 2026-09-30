package cgosotest

/*
// intentionally write the same LDFLAGS differently
// to test correct handling of LDFLAGS.
#cgo linux LDFLAGS: -L. -lcgosotest
#cgo dragonfly LDFLAGS: -L. -l cgosotest
#cgo freebsd LDFLAGS: -L. -l cgosotest
#cgo openbsd LDFLAGS: -L. -l cgosotest
#cgo solaris LDFLAGS: -L. -lcgosotest
#cgo netbsd LDFLAGS: -L. libcgosotest.so
#cgo darwin LDFLAGS: -L. libcgosotest.dylib
#cgo windows LDFLAGS: -L. libcgosotest.a
#cgo aix LDFLAGS: -L. -l cgosotest

void init(void);
void sofunc(void);
*/
import "C"

func Test() {
	C.init()
	C.sofunc()
}

//export goCallback
func goCallback() {
}
