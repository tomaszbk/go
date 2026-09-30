//go:build ignore

package main

/*
struct Issue38649 { int x; };
#define issue38649 struct Issue38649
*/
import "C"

type issue38649 C.issue38649
