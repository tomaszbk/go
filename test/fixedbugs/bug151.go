// compile


package bug151

type S string

type Empty interface {}

func (v S) Less(e Empty) bool {
	return v < e.(S);
}

/*
bugs/bug151.go:10: illegal types for operand: CALL
	string
	S
*/
