// compile


package P

const a = 0;

func f(a int) {
	a = 0;
}

/*
bug161.go:8: operation LITERAL not allowed in assignment context
*/
