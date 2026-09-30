//go:build ignore

package main

/*
// from <linux/kcm.h>
struct issue48396 {
	int fd;
	int bpf_fd;
};
*/
import "C"

type Issue48396 C.struct_issue48396
