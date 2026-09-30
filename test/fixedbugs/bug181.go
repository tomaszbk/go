// errorcheck


package main

type T *struct {
	T;	// ERROR "embed.*pointer"
}
