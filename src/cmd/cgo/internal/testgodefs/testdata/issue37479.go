//go:build ignore

package main

/*
typedef struct A A;

typedef struct {
	struct A *next;
	struct A **prev;
} N;

struct A
{
	N n;
};

typedef struct B
{
	A* a;
} B;
*/
import "C"

type N C.N

type A C.A

type B C.B
