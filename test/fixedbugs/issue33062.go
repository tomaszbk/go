// run


// Issue 33062: gccgo generates incorrect type equality
// functions.

package main

type simpleStruct struct {
	int
	string
}

type complexStruct struct {
	int
	simpleStruct
}

func main() {
	x := complexStruct{1, simpleStruct{2, "xxx"}}
	ix := interface{}(x)
	y := complexStruct{1, simpleStruct{2, "yyy"}}
	iy := interface{}(y)
	if ix != ix {
		panic("FAIL")
	}
	if ix == iy {
		panic("FAIL")
	}
}
