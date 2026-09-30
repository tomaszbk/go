// compile


package bug150

type T int
func (t T) M()

type M interface { M() } 

func g() (T, T)

func f() (a, b M) {
	a, b = g();
	return;
}

/*
bugs/bug150.go:13: reorder2: too many function calls evaluating parameters
*/
