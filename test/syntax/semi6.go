// errorcheck


package main

type T1	// ERROR "newline in type declaration"

type T2 /* // ERROR "(semicolon.*|EOF) in type declaration" */