// run


package main

type Element interface {
}

type Vector struct {
}

func (v *Vector) Insert(i int, e Element) {
}


func main() {
	type I struct { val int; };  // BUG: can't be local; works if global
	v := new(Vector);
	v.Insert(0, new(I));
}
/*
check: main_sigs_I: not defined
*/
