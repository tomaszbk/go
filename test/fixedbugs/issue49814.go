// errorcheck


package main

// "must be integer" error is for 32-bit architectures
type V [1 << 50]byte // ERROR "larger than address space|invalid array length"

var X [1 << 50]byte // ERROR "larger than address space|invalid array length"

func main() {}
