// errorcheck


package main

func f() {
	a := true;
	a |= a;	// ERROR "illegal.*OR|bool|expected"
}
