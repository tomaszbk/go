// run


package main

func main() {
	c := 10;
	d := 7;
	var x [10]int;
	i := 0;
	/* this works:
	q := c/d;
	x[i] = q;
	*/
	// this doesn't:
	x[i] = c/d;	// BUG segmentation fault
}
