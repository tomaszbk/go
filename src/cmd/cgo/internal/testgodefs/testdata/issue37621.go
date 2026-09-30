//go:build ignore

package main

/*
struct tt {
	long long a;
	long long b;
};

struct s {
	struct tt ts[3];
};
*/
import "C"

type TT C.struct_tt

type S C.struct_s
