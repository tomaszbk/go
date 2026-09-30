// errorcheck


package main

type T int

func f() {
	var x struct { T };
	var y struct { T T };
	x = y;	// ERROR "cannot|incompatible"
	_ = x;
}

type T1 struct { T }
type T2 struct { T T }

func g() {
	var x T1;
	var y T2;
	x = y;	// ERROR "cannot|incompatible"
	_ = x;
}

