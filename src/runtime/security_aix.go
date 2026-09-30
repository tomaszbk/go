package runtime

// secureMode is only ever mutated in schedinit, so we don't need to worry about
// synchronization primitives.
var secureMode bool

func initSecureMode() {
	secureMode = !(getuid() == geteuid() && getgid() == getegid())
}

func isSecureMode() bool {
	return secureMode
}
