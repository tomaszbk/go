package main

// #include <complex.h>
import "C"

//export exportComplex64
func exportComplex64(v complex64) complex64 {
	return v
}

//export exportComplex128
func exportComplex128(v complex128) complex128 {
	return v
}

//export exportComplexfloat
func exportComplexfloat(v C.complexfloat) C.complexfloat {
	return v
}

//export exportComplexdouble
func exportComplexdouble(v C.complexdouble) C.complexdouble {
	return v
}

func main() {}
