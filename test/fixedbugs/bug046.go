// errorcheck


package main

type T *struct {}

func (x T) M () {}  // ERROR "pointer|receiver"

/*
bug046.go:7: illegal <this> pointer
*/
