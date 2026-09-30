package cgotest

// typedef struct { } T42495A;
// typedef struct { int x[0]; } T42495B;
import "C"

//export Issue42495A
func Issue42495A(C.T42495A) {}

//export Issue42495B
func Issue42495B(C.T42495B) {}
