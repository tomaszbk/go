// errorcheck


package main

type T func()

type I interface {
	f, g ();  // ERROR "unexpected comma"
}

type J interface {
	h T;  // ERROR "syntax|signature"
}
