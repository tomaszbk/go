//go:build !linux && !darwin && !dragonfly && !freebsd && !netbsd && !solaris

package runtime

func sysargs(argc int32, argv **byte) {
}
