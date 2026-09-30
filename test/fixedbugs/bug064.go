// compile


package main

func
swap(x, y int) (u, v int) {
	return y, x
}

func
main() {
	a := 1;
	b := 2;
	a, b = swap(swap(a, b));
	if a != 2 || b != 1 {
		panic("bad swap");
	}
}
