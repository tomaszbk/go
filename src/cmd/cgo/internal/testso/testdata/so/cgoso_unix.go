//go:build aix || dragonfly || freebsd || linux || netbsd || solaris

package cgosotest

/*
extern int __thread tlsvar;
int *getTLS() { return &tlsvar; }
*/
import "C"

func init() {
	if v := *C.getTLS(); v != 12345 {
		println("got", v)
		panic("BAD TLS value")
	}
}
