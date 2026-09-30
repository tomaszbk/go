// errorcheck


package main

func f() int { }	// ERROR "return|control"
func g() (foo int) { }	// ERROR "return|control"
