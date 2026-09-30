// run


// Test simple multi-argument multi-valued function.

package main

func
main() {
	var x,y int;

	x,y = simple(10,20,30);
	if x+y != 65 { panic(x+y); }
}

func
simple(ia,ib,ic int) (oa,ob int) {
	return ia+5, ib+ic;
}
