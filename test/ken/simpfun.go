// run


// Test simple functions.

package main

func
main() {
	var x int;

	x = fun(10,20,30);
	if x != 60 { panic(x); }
}

func
fun(ia,ib,ic int)int {
	var o int;

	o = ia+ib+ic;
	if o != 60 { panic(o); }
	return o;
}
