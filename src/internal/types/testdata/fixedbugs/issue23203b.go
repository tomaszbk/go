package main

import "unsafe"

type T struct{}

func (T) m2([unsafe.Sizeof(T.m1)]int) {}
func (T) m1()                         {}

func main() {}
