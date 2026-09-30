//go:build ignore

package main

// enum { ENUMVAL = 0x1 };
import "C"

const ENUMVAL = C.ENUMVAL
