// run


package main

func main() {
	s := 0;
	for _, v := range []int{1} {
		s += v;
	}
	if s != 1 {
		println("BUG: s =", s);
	}
}
