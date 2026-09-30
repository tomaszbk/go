package a

var pl int

type NoitfStruct struct {
	F int
	G int
}

//go:nointerface
func (t *NoitfStruct) NoInterfaceMethod() {}
