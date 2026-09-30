package runtime

import _ "unsafe" // for go:linkname

//go:linkname setMaxStack runtime/debug.setMaxStack
func setMaxStack(in int) (out int) {
	out = int(maxstacksize)
	maxstacksize = uintptr(in)
	return out
}

//go:linkname setPanicOnFault runtime/debug.setPanicOnFault
func setPanicOnFault(new bool) (old bool) {
	gp := getg()
	old = gp.paniconfault
	gp.paniconfault = new
	return old
}
