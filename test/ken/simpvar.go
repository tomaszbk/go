// run


// Test scoping of variables.


package main

var	x,y	int;

func
main() {

	x = 15;
	y = 20;
	{
		var x int;
		x = 25;
		y = 25;
		_ = x;
	}
	x = x+y;
	if(x != 40) { panic(x); }
}
