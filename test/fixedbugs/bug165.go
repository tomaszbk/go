// errorcheck


package main

type I interface {
	m(map[I] bool);	// ok
}

type S struct {
	m map[S] bool;	// ERROR "map key type"
}
