package main

import "testplugin/iface_i"

//go:noinline
func F() interface{} {
	return (*iface_i.T)(nil)
}

//go:noinline
func G() iface_i.I {
	return (*iface_i.T)(nil)
}
